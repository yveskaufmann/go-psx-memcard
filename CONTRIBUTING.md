# Contributing to psx-memcard

## Prerequisites

- Go 1.24+ installed and configured.
- Pull latest `main` before starting work.

## Development strategy

- Trunk-based: short-lived feature branches off `main`.
- Keep changes small and reviewable; stack PRs only when necessary.

## Branch naming

- Pattern: `type/<ticket-or-id>/short-description` (e.g., `feat/PROJ-10/ui-fix`).
- Types and examples:
	- feat: `feat/123-add-block-copy`
	- fix: `fix/456-null-pointer`
	- chore: `chore/789-update-deps`
	- refactor: `refactor/321-simplify-reader`
	- docs: `docs/654-update-readme`
	- perf: `perf/987-speed-up-icon-load`

## Commit messages

- Semantic format: `type(scope): short description` (e.g., `feat(ui): add block copy`).

## Before opening a PR

- Format code (`gofmt`/`go fmt ./...`).
- Tests pass: `go test ./...`.
- Rebase on latest `main`; resolve conflicts.
- Write a clear PR description (problem, solution, focus areas, related tickets).

## Merging

- Prefer squash-and-merge for routine work.
- Use merge commits for large architectural/epic changes where history matters.

## Releases

- SemVer examples:
	- Major: remove or rename exported APIs; change card file format in a breaking way.
	- Minor: add a new UI view or block copy feature; improve icon rendering performance.
	- Patch: fix a panic in card reader; dependency security update; documentation-only change.
- Tagging: `git tag vX.Y.Z -m "release vX.Y.Z"` then `git push origin vX.Y.Z`.

## Feature guidance

- UI features: add under `internal/ui/` following view/view-model pattern; keep logic in view-models.
- Domain logic: extend `internal/memcard/`; register new services in `internal/dig/container.go`.
- Tests: table-driven where appropriate; keep isolated in `*_test.go`.

## References

- [Agents Guide](.github/agents.md)
- [Copilot Instructions](.github/copilot-instructions.md)
- [SemVer](https://semver.org/)
