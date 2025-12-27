# Agents Guide

Use these when automation (agents) creates branches, PRs, or releases.

## Branching

- Strategy: trunk-based; short-lived feature branches off `main`.
- Naming: `type/<ticket-or-id>/short-description` (e.g., `feat/PROJ-10/ui-fix`).
- Prefixes and examples:
  - feat: `feat/123-add-block-copy`
  - fix: `fix/456-null-pointer`
  - chore: `chore/789-update-deps`
  - refactor: `refactor/321-simplify-reader`
  - docs: `docs/654-update-readme`
  - perf: `perf/987-speed-up-icon-load`

## Commits

- Semantic format: `type(scope): short description` (e.g., `feat(ui): add block copy`).
- Keep commits small and focused; prefer squash on merge unless history is needed.

## PRs

- PR title: concise problem/solution.
- Description should cover: problem, solution, attention areas, related tickets.
- Checklist before opening: code formatted, tests pass (`go test ./...`), branch rebased on `main`.
- Respond promptly to review feedback; split complex follow-ups if needed.

## Releases

- Versioning: SemVer (major/minor/patch) based on change impact.
- Tagging: `git tag vX.Y.Z -m "release vX.Y.Z"` then `git push origin vX.Y.Z`.
- Use squash-and-merge for routine changes; use merge commits for large architectural work when history matters.
- SemVer examples:
  - Major: remove/rename exported APIs, change saved card format in a breaking way.
  - Minor: add block copy feature, improve performance of icon rendering, add new UI view.
  - Patch: fix panic in reader, dependency security update, documentation-only change.

## Quality gates

- Tests: must pass `go test ./...`.
- Formatting: ensure `gofmt` (or `go fmt ./...`) clean.
- Register new dependencies in `internal/dig/container.go` when adding services.

## References

- [SemVer](https://semver.org/)
- Internal: [CONTRIBUTING.md](../CONTRIBUTING.md)
- Internal: [Copilot Instructions](copilot-instructions.md)

## Documentation Lookup (Context7)

- Always fetch latest library/SDK docs via Context7 when uncertain.
- Steps:
  - Resolve the library ID by name (e.g., "fyne").
  - Fetch docs with a focused topic (e.g., "widgets", "binding", "routing").
  - Use a reasonable token budget (e.g., 2000–5000) for complex topics.
- Guidance:
  - Prefer authoritative matches and concise snippets; avoid large verbatim copies.
  - Adapt examples to project conventions (MVVM, DI, tests).

## Project-wide MCP configuration

- Responsibility: Humans configure MCP servers and provide API keys/secrets.
- Agents must not create, edit, or commit MCP configurations or secrets.
- Safety: Allowlist read-only tools where possible; avoid `*` unless necessary.
