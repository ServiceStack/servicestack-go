#!/usr/bin/env node
// Updates the version in every file that declares it, then commits the change.
//
//   npm run bump            # 0.1.2 -> 0.1.3
//   npm run bump -- minor   # 0.1.2 -> 0.2.0
//   npm run bump -- major   # 0.1.2 -> 1.0.0
//   npm run bump -- 1.2.3   # explicit version
import { readFileSync, writeFileSync } from 'node:fs'
import { execFileSync } from 'node:child_process'
import { versionFiles, replacements } from './release.config.mjs'

const run = (cmd, args) => execFileSync(cmd, args, { encoding: 'utf8' }).trim()

const pkg = JSON.parse(readFileSync('package.json', 'utf8'))
const current = pkg.version

const arg = process.argv[2] ?? 'patch'
const next = (() => {
    if (/^\d+\.\d+\.\d+$/.test(arg)) return arg
    const [major, minor, patch] = current.split('.').map(Number)
    if (arg === 'major') return `${major + 1}.0.0`
    if (arg === 'minor') return `${major}.${minor + 1}.0`
    if (arg === 'patch') return `${major}.${minor}.${patch + 1}`
    console.error(`Usage: npm run bump [-- patch|minor|major|x.y.z], got '${arg}'`)
    process.exit(1)
})()

if (next === current) {
    console.error(`Already at v${current}`)
    process.exit(1)
}

// package.json is the source of truth the other manifests are kept in sync with
pkg.version = next
writeFileSync('package.json', `${JSON.stringify(pkg, null, 2)}\n`)

for (const { file, find, replace } of versionFiles) {
    const src = readFileSync(file, 'utf8')
    const from = find(current)
    if (!src.includes(from)) {
        console.error(`${file}: could not find '${from}'`)
        process.exit(1)
    }
    writeFileSync(file, src.replace(from, replace(next)))
    console.log(`Updated ${file}`)
}

for (const { file, find, replace } of replacements ?? []) {
    const src = readFileSync(file, 'utf8')
    const updated = src.replaceAll(find(current), replace(next))
    if (updated !== src) {
        writeFileSync(file, updated)
        console.log(`Updated ${file}`)
    }
}

// Start a CHANGELOG entry for the new version when it doesn't have one yet
const changelog = readFileSync('CHANGELOG.md', 'utf8')
if (!changelog.includes(`## [${next}]`)) {
    const heading = changelog.indexOf('## [')
    const entry = `## [${next}]\n\n### Added\n\n- TODO: describe this release\n\n`
    const withEntry = heading === -1
        ? `${changelog.trimEnd()}\n\n${entry}`
        : changelog.slice(0, heading) + entry + changelog.slice(heading)
    const link = `[${next}]: https://github.com/${pkg.repository}/releases/tag/v${next}`
    writeFileSync('CHANGELOG.md', withEntry.includes(link)
        ? withEntry
        : withEntry.replace(/\n(\[\d+\.\d+\.\d+\]: )/, `\n${link}\n$1`))
    console.log('Updated CHANGELOG.md — describe the release before running `npm run release`')
}

run('git', ['add', '-A'])
run('git', ['commit', '-m', `v${next}`])
console.log(`\nBumped v${current} -> v${next}`)
