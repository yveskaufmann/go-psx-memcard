# Copilot Instructions for go-psx-memcard

## Project Overview

This is a Go application for managing PlayStation memory card files, with a GUI built using Fyne. The codebase is organized into clear internal modules for memory card logic and UI components.

## Architecture

- **Main entry point:** `main.go` initializes the app and wires up UI and logic.
- **Internal modules:**
  - `internal/dig/`: Container logic (dependency injection, service management).
  - `internal/memcard/`: Core memory card data structures and operations (block management, reading, encoding, icon handling).
  - `internal/ui/`: Fyne-based UI, split into subcomponents:
    - `manager-view*`: Main card manager view and its view-model.
    - `blocks/`: Block grid and block views.
    - `filepicker/`: File picker dialog and logic.
    - `utils/`: Window utilities.
    - `animated-sprite/`: Sprite animation for icons.
  - The UI components follow a view/view-model pattern for separation of concerns.
    - Views are responsible for rendering and are also the orchestration point for Fyne widgets.
    - State and logic are handled in the corresponding view-models.
    - View-Models interact with the model in internal/memcard/ to perform operations and update state.
    - View-Models do not directly manipulate Fyne widgets; instead, they expose state and methods that the views bind to.View models do not directly manipulate views; they should not have direct access to views or widgets.
    - views bind to view-model properties and methods to update the UI.
    - views can also register to event listeners from the view-model to react to changes or actions.

## Developer Workflows

- **Build:** Use `make` or run `go build` from the project root. The output binary is placed in `bin/psx-memcard`.
- **Run:** Execute the binary directly or use `make run` if defined.
- **Test:** Unit tests are in files ending with `_test.go` (e.g., `sjis-string_test.go`). Run `go test ./...` from the root.
- **Debug:** UI logic is in `internal/ui/`; memory card logic is in `internal/memcard/`. Use Go's standard debugging tools.

## Conventions & Patterns

- **Dependency injection:** Managed via `internal/dig/container.go`.
- **UI pattern:** Fyne widgets are wrapped in view and view-model pairs for separation of concerns.
- **Memory card abstraction:** All block and card operations are encapsulated in `internal/memcard/`.
- **File structure:** Each UI feature is split into its own directory for maintainability.
- **No global state:** State is passed via view-models and containers.

## Integration Points

- **External dependencies:** Fyne (GUI), Go standard library.
- **File I/O:** Memory card files are loaded from `dummy-cards/` for testing and demo purposes.

## Examples

- To add a new UI feature, create a new directory under `internal/ui/`, following the view/view-model pattern.
- To extend memory card logic, add new methods to the appropriate file in `internal/memcard/`.
