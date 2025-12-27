# Overview of psx-memcard Application

This is a application for managing PlayStation memory card image files, written in Go.
PlayStation memory card images are used by emulators to persist save games of playstations.
It is intended to help users of emulators manage their memory card files like in the original PlayStation interface, but with additional features and a modern GUI.

This app allows users to view, create, manipulate playstation memory card images by supporting up to two memory cards loaded simultaneously, allowing users to:

- view their contents side-by-side Key:
- each memory cards blocks are rendered in a grid similar to the original PlayStation interface
- users can see detailed information about each block, including animated icons and metadata
- copy blocks between the two loaded memory cards
- create new memory card files from scratch by opening a new file via a file picker dialog
- delete existing blocks from a memory card
- changes are reflected immediately in the UI, with visual feedback for operations like deletion and copying

## Technology Stack

- **Languages**: Golang, Bash (Makefile)
- **GUI Framework**: Fyne
- **Build System**: Go Modules
- **Testing**: Go's built-in testing framework
- **Formatting**: gofmt + editorconfig; use `gofmt` to format code consistently.
- **Dependency Management**: Go Modules
- **Dependency injection**: Managed via `internal/dig/container.go`.
- **Build System**: Makefile for automate common tasks (build, run, test)

## Project and code guidelines

- Follow Go best practices and idioms.
- Write GoDoc comments for all exported functions, structs, and packages.
  - Ensure GoDoc enables usage without reading source; include examples when helpful.
  - Start each GoDoc comment with the entity name (e.g., "FuncName does X...").
  - Describe how to use the entity (what and why), not its internal implementation, unless implementation details affect usage.
  - Ensure GoDoc enables usage without reading source; include examples when helpful.
- Write unit tests for all non-trivial functions and methods.
  - Use descriptive names for test functions (e.g., TestCalculateChecksum_ValidInput).
  - Cover edge cases and error conditions.
  - Keep tests isolated; avoid dependencies on external systems or state.
  - Use table-driven tests for functions with multiple input scenarios.
  - Avoid to comment out failing tests; fix the underlying issue instead.
  - Avoid to change tests to match broken code; fix the broken code instead; only change tests when the requirements changes.
- Keep functions small and focused (ideally under 50 lines).
- Use meaningful names for variables, functions, structs, and packages.
- Use depencency injection via constructors or function parameters; avoid global state.
  - prefer explicit dependencies passed via constructor functions
  - bind dependencies within the main DI container in `internal/dig/container.go`
  - consider module level containers for complex modules and inject them into the main container
  - Avoid circular dependencies between packages; use interfaces to break cycles if necessary.
- Use idiomatic Go error handling (return errors, not panic) - unless within in the main package or in tests.
- Follow consistent formatting using gofmt and editorconfig.
- Organize the codebase into clear internal modules for memory card logic and UI components; keep related code together; separate concerns e.g UI vs domain / application logic.
- Use Makefile to automate common tasks (build, run, test)
  - Explain each goal target with comments in the Makefile.
  - Keep goal recipes simple and short; avoid complex shell logic in the Makefile instead use separete scripts in the scripts/ folder for complex logic.
  - keep build scripts within the scripts/ folder for complex logic.
  - Build scripts should be written in bash for portability
    - avoid using platform-specific tools; if not possible document platform dependencies clearly within in the README.md
    - aim for POSIX compliance where possible.
    - use `set -euo pipefail` at the start of scripts to ensure robust error handling.
    - use functions within scripts to encapsulate reusable logic.
    - document each script's purpose and usage at the top of the file.
    - use clear and descriptive names for scripts and functions.
    - apply single responsibility principle to keep script functions focused on one concern.
- To add a new UI feature, create a new directory under `internal/ui/`, following the view/view-model pattern.
- To extend memory card logic, add new methods to the appropriate file in `internal/memcard/`.

### UI Guidelines

- **Framework:** Fyne for cross-platform GUI.
- **UI pattern:** Fyne widgets are wrapped in view and view-model pairs for separation of concerns.
- **File structure:** Each UI feature is split into its own directory for maintainability.
- **No global state:** State is passed via view-models and containers.
- **Architecture:** MVVM-inspired pattern with view and view-model separation; UI components follow the view/view-model pattern:
  - Views are responsible for rendering and are also the orchestration point for Fyne widgets.
  - State and logic are handled in the corresponding view-models.
  - View-Models interact with the model in internal/memcard/ to perform operations and update state.
  - View-Models do not directly manipulate Fyne widgets; instead, they expose state and methods that the views bind to.View models do not directly manipulate views; they should not have direct access to views or widgets.
  - views bind to view-model properties and methods to update the UI.
  - views can also register to event listeners from the view-model to react to changes or actions.

## Project Structure

- Root files:

  - `main.go`: application entry point; initializes DI container and starts the Fyne UI.
  - `go.mod`/`go.sum`: Go module definitions.
  - `Makefile`: build, run, and test shortcuts (e.g., `make`, `make run`).
  - `FyneApp.toml`: Fyne app metadata.
  - `dummy-cards/`: sample memory card images for testing (`epsxe000.mcr`, `epsxe001.mcr`).
  - `bin/`: built binaries (e.g., `psx-memcard`).

- `internal/` (primary implementation):

  - `dig/`: dependency-injection container and wiring (`container.go`).
  - `memcard/`: core domain code for memory cards and blocks (`block.go`, `data.go`, `read.go`, `icon.go`, `block-mgnt.go`, `sjis-string.go` and tests).
  - `ui/`: Fyne UI components and view-models:
    - manager view & view-model: `manager-view.go`, `manager-view-model.go`, `ui.go`.
    - `blocks/`: block grid, block views and view-models (`blocks-container.go`, `block-view.go`, `block-view-model.go`, `selection-model.go`).
    - `filepicker/`: file-picker view and view-model.
    - `animated-sprite/`: sprite animation utilities.
    - `utils/`: window and UI helper utilities.

- Tests and docs:

  - Unit tests live next to packages (`*_test.go`, e.g., `sjis-string_test.go`).
  - Write GoDoc comments for exported entities and keep tests isolated.

- Guidance for contributing:
  - Add UI features under `internal/ui/` following the view/view-model pattern.
  - Keep domain logic in `internal/memcard/` and register dependencies in `internal/dig/container.go`.

## Resources

- scripts folders: `scripts/` for automation scripts (e.g., release scripts)
- Makefile: for build tasks (e.g.:build, run, test)
- FyneApp.toml: Fyne app metadata
- Dummy memory card images: `dummy-cards/epsxe000.mcr`, `example cards for testing and demo purposes.

### Developer Workflow Scripts

- **Build:** Use `make` or run `go build` from the project root. The output binary is placed in `bin/psx-memcard`.
- **Run:** Execute the binary directly or use `make run` if defined.
- **Test:** Unit tests are in files ending with `_test.go` (e.g., `sjis-string_test.go`). Run `go test ./...` from the root.
- **Debug:** UI logic is in `internal/ui/`; memory card logic is in `internal/memcard/`. Use Go's standard debugging tools.

### Useful Links

- [Fyne Documentation](https://developer.fyne.io/)
- [Go Documentation](https://golang.org/doc/)
- [PlayStation Memory Card Format](https://problemkaputt.de/psx-spx.htm#memorycardfileformat)

### Engineering Workflow & Contribution Guide

#### 1. Environment & Setup

- **Prerequisites:** Ensure **Go > 1.24** is installed and correctly configured.
- **Repository:** Clone or fork the repository. If already cloned, always `git pull` the latest changes from `main` before starting work.

---

#### 2. Development Strategy

We follow **Trunk-Based Development** to ensure high velocity and integration frequency.

- **Branching:** Create short-lived feature branches off the `main` branch.
- **Small PRs:** Split large changes into small, reviewable Pull Requests.
- **Stacked PRs:** If a feature depends on unmerged code from a previous task, use stacked PRs.
  - _Branch A_ → Base: `main`
  - _Branch B_ → Base: _Branch A_

---

#### 3. Branch Naming Convention

Follow the pattern: `type/[ticket-id]/short-description`

| Prefix        | Purpose                                 | Example                      |
| :------------ | :-------------------------------------- | :--------------------------- |
| **feat/**     | A new feature or functionality          | `feat/PROJ-10/user-auth`     |
| **fix/**      | A bug fix                               | `fix/PROJ-22/login-flicker`  |
| **chore/**    | Maintenance, dependencies, or tooling   | `chore/update-node-version`  |
| **refactor/** | Code change (no bug fix or new feature) | `refactor/api-client`        |
| **docs/**     | Documentation changes only              | `docs/api-readme`            |
| **perf/**     | A code change that improves performance | `perf/db-query-optimization` |

---

#### 4. Commit Message Standard

We use **Semantic Commits** to keep the history searchable and meaningful.  
**Format:** `type(scope): short description`

- **Example:** `feat(ui): add dark mode support`
- **Example:** `fix(api): resolve null pointer in user-service`

---

#### 5. Before Opening a Pull Request

To respect the reviewers' time and maintain code quality, complete this checklist:

1.  **Validation:** Ensure all tests pass and code is formatted (`go fmt`).
2.  **Clean History:** Use `git rebase -i` to consolidate your work into a few clean, atomic commits categorized by scope (e.g., `feat(logic)` and `docs(api)`) to ensure the history is both readable and easy to analyze.
3.  **Sync:** Use `git pull --rebase` to sync with the `main` branch and resolve any conflicts.
4.  **Standards:** Verify changes adhere to the project's coding standards.

---

#### 6. Creating & Managing Pull Requests (PR)

##### PR Description

Explain the changes clearly by addressing:

- **The Problem:** What is being solved?
- **The Solution:** Describe the key changes made.
- **Attention Areas:** Highlight specific areas for reviewers to focus on.
- **Traceability:** Reference any related GitHub ticket numbers.

##### The Review Process

- **Promptness:** Address requested changes promptly.
- **Complexity:** If requested changes are complex, discuss with reviewers whether to break them into smaller follow-up issues.
- **Iteration:** Maintain a balance between perfection and iteration; when in doubt, prefer smaller, incremental changes.

---

#### 7. Merging & Cleanup

We use a dual-merge strategy based on the impact of the change:

- **Primary Strategy (Squash and Merge):** Preferred for standard features, bug fixes, and chores. This keeps the `main` branch history clean, concise, and easy to revert.
- **Architectural Strategy (Merge Commit):** Reserved for large architectural changes or "Epics" where the granular evolution of the code (the "how" and "why") is important for future analysis and forensics.
- **Housekeeping:** After merging, **delete the feature branch** (both locally and on GitHub) to keep the repository clean.

---

#### 8. Creating Releases

We follow **Semantic Versioning (SemVer)** for all project releases. Releases are triggered by creating a git tag following the pattern `v<semantic-version>`.

##### Versioning Logic

Choose the next version number based on the nature of the changes since the last tag:

- **Major (vX.0.0):** Increment for breaking changes that signal compatibility shifts.
- **Minor (v0.X.0):** Increment for new features, performance improvements, or UI changes.
- **Patch (v0.0.X):** Increment for bug fixes, security patches, documentation updates, or chore tasks.

##### Tagging Procedure

To create a release, tag the desired commit on the `main` branch:

```bash
git tag v1.2.3 -m "release v1.2.3"
git push origin v1.2.3
```

#### Guidance for Adding new Features

Additional guidelines for contributing new features or extending functionality:

1. Create a new directory under `internal/ui/` for the feature.
2. Implement the view and view-model following the established pattern.
3. Register any new dependencies in `internal/dig/container.go`.
4. Write unit tests for the view-model logic.
5. Ensure the UI components are properly integrated into the main application.

##### Extending Memory Card Logic

1. Identify the appropriate file in `internal/memcard/` for the new functionality.
2. Add new methods or structs as needed.
3. Update the DI container in `internal/dig/container.go` if new dependencies are introduced.
4. Write unit tests to cover the new logic.

## Domain Concepts

- **Memory Card:** Represents a PlayStation memory card image file, containing multiple blocks.
- **Block:** A fixed-size unit of data on the memory card, which can store save game data or be empty.
- **Icon:** Visual representation of a saved game, including animated sprites
- **Memory Card Operations:** Actions that can be performed on memory cards, such as loading, saving,
  copying blocks, and deleting blocks. Kept in `internal/memcard/`.

```

```
