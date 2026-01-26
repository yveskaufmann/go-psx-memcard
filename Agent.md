# Agent.md — Repository agent instructions

This file is the canonical agent instruction document for this repository. It follows the open Agents standard and preserves the guidance previously stored in `.github/copilot-instructions.md` and `.github/agents.md`.

Purpose
- Provide an authoritative, deduplicated set of instructions for automated agents and human collaborators interacting with this repository.
- Preserve existing rules and policies while adopting the [agents.md](https://agents.md/) convention.

Scope
- Applies to automation that creates branches, pull requests, releases, or performs repository edits.

Core guidelines (merged and deduplicated)

1) Project summary
- Go application to view and manage PlayStation memory card images (`.mcr`) with a Fyne GUI.
- Key paths: `main.go`, `internal/` (domain and UI), `internal/dig/container.go`, `Makefile`, `FyneApp.toml`, `dummy-cards/` (samples).

2) Development & tooling
- Go 1.24+; formatting with `gofmt` / `go fmt`.
- Tests: `go test ./...` — unit tests live alongside code as `*_test.go` files.
- Script automation via `Makefile` and `scripts/` (use Makefile targets where provided).

3) Style & code rules
- GoDoc: Comment exported identifiers; start comments with the name being documented and include brief examples where helpful.
- Tests: Use table-driven tests where appropriate; cover edge/error cases and keep tests isolated.
- Dependency injection: Prefer constructor injection; register services in `internal/dig/container.go`.
- Error handling: Return errors; only `panic` in `main` or tests.
- Shell scripts: Use `set -euo pipefail` and encapsulate logic in functions.

4) UI & architecture conventions
- MVVM pattern: Views (Fyne) render UI; ViewModels hold state and coordinate with domain logic (`internal/memcard/`).
- ViewModels must not import or manipulate view widgets directly; use bindings and events.

5) Repo workflow (branches / commits / PRs / releases)
- Branching: trunk-based — short-lived feature branches off `main`.
- Branch naming: `type/<ticket-or-id>/short-description` (e.g., `feat/123-add-block-copy`).
- Commit style: `type(scope): short description` (e.g., `feat(ui): add block copy`).
- PRs: Title = concise problem/solution. PR description should cover problem, solution, attention areas and related tickets. Checklist before opening PR: code formatted, tests pass, branch rebased on `main`.
- Releases: Use SemVer. Tag with `git tag vX.Y.Z -m "release vX.Y.Z"` then `git push origin vX.Y.Z`.

6) Quality gates
- CI must run `gofmt` (or fail on formatting differences), `go vet`, `go test` and `make build` for verification.

7) External docs & knowledge lookup
- Use an external documentation service (Context7) only when available: resolve library IDs then fetch docs by topic (for example: `fyne` → `widgets`, `binding`). Prefer concise excerpts and adapt examples to project MVVM and DI patterns. Do not create or edit MCP configs; humans manage keys.

8) Security and secrets
- Do not commit secrets or configuration keys. For any required secret usage, ask a human to add them to repository secret storage.

Agent behavior & non-regression
- Agents must act according to this `Agent.md` policy. To preserve existing behavior we keep the precise rules from the archived originals (see `docs/agents-orig/`).

Refresh context skill (recommended)
- On new session start, agents SHOULD read and load the repository architecture artifacts to establish context.
- Required initial reads:
  - `architecture.md`
  - `architecture.puml`
- Provide a `refreshContext()` skill or command that re-reads those files and updates the agent session context when requested. Example instruction for an agent:

  "At session start, run `readFile('architecture.md')` and `readFile('architecture.puml')` to populate architectural context. Expose `refreshContext()` that repeats these reads on demand. Use this to reconcile long-running agents with repo changes."

Attribution & preservation
- This document consolidates content from the following originals (preserved verbatim in `docs/agents-orig/`):
  - `.github/copilot-instructions.md`
  - `.github/agents.md`

Usage notes
- Keep `docs/agents-orig/` as an audit trail. After review and acceptance, `.github/copilot-instructions.md` will be removed to avoid duplication; the canonical instructions will live in `Agent.md`.

If you need a stricter machine-readable schema of these rules for enforcement, add a small `scripts/validate-agent-md.sh` that checks presence of required sections and that `architecture.md` is referenced.