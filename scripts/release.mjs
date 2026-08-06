#!/usr/bin/env node
// Tags the current version and creates the GitHub Release that triggers the
// publish workflow.
//
//   npm run release            # release the current version
//   npm run release -- patch   # bump first, then release (also minor|major|x.y.z)
import { readFileSync } from 'node:fs'
import { execFileSync } from 'node:child_process'
import { versionFiles } from './release.config.mjs'

const run = (cmd, args, opts = {}) =>
    execFileSync(cmd, args, { encoding: 'utf8', stdio: opts.inherit ? 'inherit' : 'pipe' })?.trim()

const fail = (message) => {
    console.error(`✗ ${message}`)
    process.exit(1)
}

const bumpArg = process.argv[2]
if (bumpArg) {
    run('node', ['scripts/bump.mjs', bumpArg], { inherit: true })
}

const pkg = JSON.parse(readFileSync('package.json', 'utf8'))
const version = pkg.version
const tag = `v${version}`

// Every manifest must agree with package.json, otherwise the publish will fail
for (const { file, find } of versionFiles) {
    if (!readFileSync(file, 'utf8').includes(find(version))) {
        fail(`${file} isn't at ${version}, run \`npm run bump\` first`)
    }
}

if (run('git', ['status', '--porcelain'])) {
    fail('Working directory has uncommitted changes')
}

const branch = run('git', ['rev-parse', '--abbrev-ref', 'HEAD'])
if (branch !== 'main') {
    fail(`Releases are cut from main, currently on ${branch}`)
}

if (run('git', ['tag', '-l', tag])) {
    fail(`${tag} already exists`)
}

// Release notes come from this version's CHANGELOG entry
const changelog = readFileSync('CHANGELOG.md', 'utf8')
const section = changelog.split(`## [${version}]`)[1]?.split('\n## [')[0]?.trim()
if (!section) {
    fail(`CHANGELOG.md has no '## [${version}]' entry`)
}
if (section.includes('TODO: describe this release')) {
    fail(`CHANGELOG.md entry for ${version} is still a TODO`)
}

console.log(`Releasing ${pkg.name} ${tag}\n`)
run('git', ['push', 'origin', branch], { inherit: true })
run('git', ['tag', '-a', tag, '-m', `${tag} - ${pkg.description ?? pkg.name}`])
run('git', ['push', 'origin', tag], { inherit: true })
run('gh', ['release', 'create', tag, '--title', tag, '--notes', section], { inherit: true })

console.log(`\n✓ Released ${tag} — the release workflow publishes it from here`)
