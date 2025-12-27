# Copilot Instructions: psx-memcard

## Overview

- Go app to view and manage PlayStation memory card images (.mcr) with a Fyne UI.
- Two cards side-by-side: inspect blocks (metadata, icons), copy/delete blocks, create new cards.

## Technology Stack

- Go (Go Modules), Fyne GUI, Bash Makefile scripts.
- Formatting: gofmt + editorconfig; tests: `go test`.
- Dependency injection: `internal/dig/container.go`.

## Project Structure

- Root: `main.go` (entry/DI bootstrap), `go.mod`/`go.sum`, `Makefile`, `FyneApp.toml`, `dummy-cards/` sample `.mcr`, `bin/` built binaries, `scripts/` automation.
- `internal/`: primary implementation.
  - `dig/`: DI container (`container.go`).
  - `memcard/`: domain logic (blocks, data, icons, reads, SJIS) and tests.
  - `ui/`: views/view-models (manager view, blocks, filepicker, animated-sprite, utils).
- Tests live next to code as `*_test.go`.

## AI Rules

- GoDoc: Comment all exported types/functions/packages; start with the identifier; explain what/why/when to use; include examples when helpful.
- Testing: Prefer table-driven tests; cover edge/error cases; keep tests isolated.
- DI: Use constructor injection; register dependencies in `internal/dig/container.go`; avoid globals and cycles (use interfaces).
- MVVM (Fyne): Views orchestrate widgets; view-models own state/logic and talk to `internal/memcard/`; view-models must not touch widgets/views directly; views bind to view-model state/events.
- Error handling: Return errors; panic only in `main` or tests.
- Scripts: Bash with `set -euo pipefail`; keep logic in functions; document usage at the top.
 - Documentation lookup (Context7):
   - Use Context7 to fetch up-to-date library/SDK docs (e.g., Fyne).
   - First resolve the library ID, then fetch docs by topic (e.g., "widgets", "binding").
   - Prefer concise excerpts over long quotes; adapt examples to project style.
   - Human-configured MCP required: Do not create or edit MCP configs; humans set up servers and provide API keys.
   - Fallback: If MCP is unavailable, rely on local docs and stable APIs.

## Commands

- Build/run: `make`, `make run`.
- Tests: `go test ./...`.

## Resources

- `scripts/`: automation scripts (build/run/test/release); prefer calling scripts via Makefile.
- `Makefile`: primary entry for build/run/test tasks.
- `dummy-cards/`: sample `.mcr` files for local testing.

## Workflow cues

- Branch prefixes (examples): `feat/123-add-block-copy`, `fix/456-null-pointer`, `docs/654-update-readme`.
- SemVer examples: major (break exported API or card format), minor (add block copy view or perf boost), patch (fix panic, dependency/security update, docs-only).

## Domain Concepts

- Memory Card: 128KB image with multiple blocks.
- Block: Fixed-size unit storing save data or empty.
- Icon: Animated sprite metadata per block.

## Useful Links

- [Fyne Documentation](https://developer.fyne.io/)
- [Go Documentation](https://golang.org/doc/)
- [PlayStation Memory Card Format](https://problemkaputt.de/psx-spx.htm#memorycardfileformat)
- [Go Doc Comments](https://go.dev/doc/comment)
- [Table-Driven Tests (Go Wiki)](https://github.com/golang/go/wiki/TableDrivenTests)
- [EditorConfig](https://editorconfig.org/)
- Internal: [CONTRIBUTING.md](../CONTRIBUTING.md)
- Internal: [Agents Guide](agents.md)
