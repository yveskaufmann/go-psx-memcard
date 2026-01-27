# Agent.md

## Purpose

Authoritative instructions for automated agents and human collaborators editing this repository.

## Scope

Applies to automation that creates branches, pull requests, releases, or performs repository edits.

## Project summary

- Go application to view and manage PlayStation memory card images (`.mcr`) with a Fyne GUI.
- Technology: Go (Go Modules), Fyne GUI, Bash/Makefile scripts.
- Key paths: `main.go`, `internal/`, `internal/dig/container.go`, `Makefile`, `FyneApp.toml`, `dummy-cards/`.

## Domain concepts

- **Memory Card**: 128KB image file containing multiple blocks.
- **Block**: Fixed-size unit storing save data or marked empty.
- **Icon**: Animated sprite metadata associated with each block.

## Development & tooling

- Target: Go 1.24+. Format with `gofmt` / `go fmt`.
- Tests: `go test ./...` (unit tests live next to code as `*_test.go`).
- Build/run: use `Makefile` targets and `scripts/` for automation.

### Makefile commands

```bash
# Build the application to bin/psx-memcard
make build

# Run unit tests
make test

# Run the application directly
make run

# Remove build artifacts and clear Go caches
make clean
```

## Style & code rules

- GoDoc: Comment exported identifiers; start comments with the identifier name.
- Tests: Prefer table-driven tests; cover edge cases and keep tests isolated.
- DI: Prefer constructor injection; register services in `internal/dig/container.go`.
- Errors: Return errors; restrict `panic` to `main` or tests.
- Shell scripts: Use `set -euo pipefail` and encapsulate logic in functions.

## UI & architecture conventions

- Follow MVVM: Views render UI; ViewModels hold state and coordinate with domain logic (`internal/memcard/`).
- ViewModels must not manipulate view widgets directly; use bindings/events.
- Refer to [`architecture.md`](architecture.md) for detailed component structure and responsibilities.
- See [`architecture.puml`](architecture.puml) for visual C4 model diagrams showing system boundaries and data flow.

### Key architectural principles

- **UI Layer** (`internal/ui/`): Fyne widgets and views (ManagerWindowView, BlocksContainer, BlockView, FilePicker, AnimatedSprite)
- **ViewModel Layer**: State management and UI coordination (ManagerWindowViewModel, BlocksContainerViewModel, SelectionViewModel)
- **Domain Layer** (`internal/memcard/`): Core business logic (MemoryCard, BlockManagement, IconDecoder, readers/writers)
- **Dependency Injection**: Use `uber/dig` container in `internal/dig/container.go`

## Repo workflow

### Prerequisites

- Go 1.24+ installed and configured.
- Pull latest `main` before starting work:

```bash
git checkout main
git pull --rebase origin main
```

### Creating a new branch

Use it when creating a new feature or bugfix branch:

```bash
# Switch to main and get latest changes
git checkout main
git pull --rebase origin main

# Create and switch to new feature branch
git checkout -b <new-branch>
```

### Rebase feature branch

Use this workflow after changes were made in the base branch or after the initial git clone:

```bash
# Pull changes from main remote and rebase current branch
git pull --rebase origin main

# Download Go dependencies
go mod download
```

### Dependency cleanup

Use when dependencies are added or removed or when imported packages change:

```bash
go mod tidy
```

### Branching

- Trunk-based; short-lived feature branches off `main`.
- Naming: `type/<ticket-or-id>/short-description`.
- Prefixes: `feat`, `fix`, `chore`, `refactor`, `docs`, `perf`.
- Examples: `feat/123-add-block-copy`, `fix/456-null-pointer`, `docs/654-update-readme`.

### Commits

- We use semantic commit messages.
- Format: `type(scope): short description`.
- Keep commits small and focused.

### PRs

- Title: concise problem/solution.
- Description: problem, solution, attention areas, related tickets.
- Checklist: code formatted, tests pass, branch rebased on `main`.

#### Small PR strategy

- **Keep PRs small**: faster review, fewer conflicts, better feedback quality.
- **Use stacked PRs for complex features**: break into logical, reviewable chunks.
- **Submit draft PRs early**: validate approach and prevent off-track development.
- **Make incremental changes**: help reviewers understand progression and maintain momentum.

### Releases

- Versioning: SemVer (major/minor/patch).
- Tag: `git tag vX.Y.Z -m "release vX.Y.Z"` then `git push origin vX.Y.Z`.
- SemVer examples:
  - **Major**: remove/rename exported APIs, break saved card format.
  - **Minor**: add block copy feature, new UI view, performance improvement.
  - **Patch**: fix panic, dependency/security update, docs-only change.

## Feature development guidance

### UI features

- Add under `internal/ui/` following view/view-model pattern.
- Keep logic in view-models; views only handle rendering and user interaction.

### Domain logic

- Extend `internal/memcard/` for business logic.
- Register new services in `internal/dig/container.go`.

### Tests

- Table-driven tests where appropriate.
- Keep tests isolated in `*_test.go` files next to source code.

## Quality gates

- CI must run `gofmt` (fail on differences), `go vet`, `go test`, and `make build`.

## External docs & knowledge lookup

- Resolve library IDs and fetch concise examples that align with project MVVM/DI patterns (e.g., `fyne` → `widgets`, `binding`).
- Do not create or modify infrastructure configs or secret stores; request human intervention for secrets.

## Security

- Never commit secrets or private keys. Ask maintainers to add required secrets to repository secret storage.

## Refresh context skill

On session start, agents SHOULD read:

- `architecture.md`
- `architecture.puml`

Use the `refresh-context` skill located in `.claude/skills/refresh-context/SKILL.md` to load architectural context on demand.

