# v0.1.0 - 2026-01-27

## Added
- Block copy feature (commit: bc55052)
- Create new memory card (commit: c954374)
- Render memory-card stats (commit: 5dc8c83)
- Block deletion (commit: bbd10ef)
- Fyne app manifest prep for packaging (commit: 0271346)
- Animated block icons (commit: 265ca83)
- App icon (commit: 3ff496b)
- CI / agent docs and workflows (commits: 00f99c9, 13bc83c)

## Fixed
- Resolve markdown formatting issues (commit: 1b53822)
- File-picker truncation fix (commit: a6a2418)
- Module path fix (commit: da7090d)

## Changed
- Documentation: architecture summary; contributing & agents guides (commits: 587156a, bf3b353)
- Refactors: MVVM pattern, memcard loading changes (commits: 8cd3b46, 85393ca)
- CI: add workflows and fix FYNE dependencies (commits: 00f99c9, f9ede06)
- Chore: stop tracking built binary, remove CLI interface, misc cleanups (commits: 47b97f, df4dfa7, ed17228)

## Notes
- This is a minor release: `v0.1.0`.
- Please run the verification steps locally before creating the release (formatting, vet, tests).

## Verification checklist (run locally)
- [ ] `go mod tidy`
- [ ] `gofmt -l .` then `gofmt -w .` to apply formatting
- [ ] `go vet ./...`
- [ ] `go test ./...`
- [ ] Commit any fixes

## Suggested command (run after verification and commit)

```bash
# create the release (gh will create the tag)
gh release create v0.1.0 --title "v0.1.0" --notes-file release.md
```
