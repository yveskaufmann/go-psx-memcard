# Architecture Documentation

## Overview

The PSX Memory Card Manager is a desktop application built with Go and the Fyne UI framework. It provides a graphical interface for managing PlayStation (PSX) memory card files (.mcr format), allowing users to view, copy, and delete save game blocks.

## System Purpose

The application enables users to:
- Load and display PlayStation memory card files
- View save game blocks with animated icons
- Copy save game blocks between memory cards
- Delete save game blocks from memory cards
- Display save game titles and metadata

## Architecture Pattern

The application follows a **Model-View-ViewModel (MVVM)** pattern with clear separation of concerns:

- **View Layer**: Fyne UI components that render the user interface
- **ViewModel Layer**: Manages UI state and coordinates between views and domain logic
- **Domain Layer**: Core business logic and data structures for PSX memory card format

## System Architecture

### High-Level Components

```
┌─────────────────────────────────────────────────────────────┐
│                    PSX Memory Card Manager                   │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │   UI Layer   │  │ ViewModel    │  │   Domain     │     │
│  │              │  │    Layer     │  │    Layer     │     │
│  │ - Views      │  │ - State Mgmt │  │ - MemoryCard │     │
│  │ - Widgets    │  │ - Commands   │  │ - I/O        │     │
│  │ - Rendering  │  │ - Bindings   │  │ - Format     │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

### Dependency Injection

The application uses **uber/dig** for dependency injection, providing a centralized container for managing component dependencies and lifecycle.

## Component Details

### 1. UI Layer (`internal/ui/`)

The UI layer is responsible for rendering the graphical interface using Fyne widgets.

#### Main Components

- **`ManagerWindowView`** (`manager-view.go`)
  - Main application window
  - Orchestrates the layout of two memory card slots (Card 1 and Card 2)
  - Contains action buttons (Copy, Delete)
  - Displays selected save game title

- **`BlocksContainer`** (`blocks/blocks-container.go`)
  - Displays memory card blocks in a 3x5 grid (15 blocks total)
  - Manages individual block views
  - Responds to data binding changes

- **`BlockView`** (`blocks/block-view.go`)
  - Renders a single memory card block
  - Shows animated save game icon
  - Handles user interaction (tap to select)
  - Visual feedback for selected/unselected state

- **`FilePicker`** (`filepicker/file-picker.go`)
  - Provides file selection dialog
  - Displays selected file path
  - Triggers memory card loading on file selection

- **`AnimatedSprite`** (`animated-sprite/sprite.go`)
  - Renders animated icons from save games
  - Supports 1, 2, or 3 frame animations
  - Handles frame timing and looping

### 2. ViewModel Layer (`internal/ui/`)

The ViewModel layer manages application state and coordinates between UI and domain logic.

#### Main Components

- **`ManagerWindowViewModel`** (`manager-view-model.go`)
  - Manages two memory card instances (left and right)
  - Handles memory card loading operations
  - Coordinates block operations (copy, delete)
  - Manages data bindings for block lists
  - Updates selected save game title

- **`BlocksContainerViewModel`** (`blocks/blocks-container-model.go`)
  - Manages state for all blocks in a memory card
  - Updates block views based on data bindings
  - Coordinates with selection model

- **`BlockViewModel`** (`blocks/block-view-model.go`)
  - Manages state for individual block
  - Handles selection state
  - Manages animation and title bindings

- **`SelectionViewModel`** (`blocks/selection-model.go`)
  - Centralized selection state management
  - Tracks selected card and block index
  - Notifies listeners of selection changes
  - Thread-safe with mutex protection

- **`FilePickerViewModel`** (`filepicker/view-model.go`)
  - Manages file path state
  - Coordinates file picker service
  - Triggers callbacks on file selection

### 3. Domain Layer (`internal/memcard/`)

The domain layer contains the core business logic and data structures for PSX memory card format.

#### Main Components

- **`MemoryCard`** (`data.go`)
  - Core data structure representing a PSX memory card
  - 128 KB total size (131,072 bytes)
  - Contains 15 data blocks (8 KB each)
  - Includes header, directory frames, and block data

- **Memory Card I/O**
  - **`Open()`** (`read.go`): Reads memory card from file
  - **`Write()`** (`write.go`): Writes memory card to file
  - Validates file size and format

- **Block Operations** (`block.go`, `block-mgnt.go`)
  - **`GetBlock()`**: Retrieves block data with title and icon
  - **`ListBlocks()`**: Lists all allocated blocks
  - **`CopyBlockTo()`**: Copies block to another memory card (TODO: implementation)
  - **`DeleteBlockFrom()`**: Deletes a block from memory card

- **Icon Decoding** (`icon.go`)
  - **`IconBitmapFrame`**: Represents 16x16 pixel icon (128 bytes)
  - **`ToImage()`**: Converts PSX icon format to Go image.Image
  - Handles 4-bit per pixel format with color palette

- **String Handling** (`sjis-string.go`)
  - **`ShiftJISString`**: Handles Shift-JIS encoding/decoding
  - Used for save game titles and file names
  - Converts between PSX format and UTF-8

#### Data Structures

- **`HeaderFrame`**: Memory card header with magic bytes and checksum
- **`DirectoryFrame`**: Directory entry for each block (allocation state, file size, file name)
- **`Block`**: Contains title frame, icon frames, and data frames
- **`BlockTitleFrame`**: Save game title and icon metadata
- **`IconBitmapFrame`**: 16x16 pixel icon bitmap data
- **`DataFrame`**: 128-byte data frame (60 frames per block)

### 4. Dependency Injection (`internal/dig/`)

- **`container.go`**: Wrapper around uber/dig
  - Provides simplified API for dependency injection
  - Manages singleton container instance
  - Handles error reporting

## Data Flow

### Loading a Memory Card

1. User selects file via `FilePicker`
2. `FilePickerViewModel` triggers callback
3. `ManagerWindowViewModel.LoadMemoryCardImage()` is called
4. `memcard.Open()` reads file from disk
5. `MemoryCard.ListBlocks()` extracts block information
6. Block data is converted to view models
7. Data bindings update UI components
8. `BlocksContainer` refreshes to show blocks

### Selecting a Block

1. User taps on a `BlockView`
2. `BlockViewModel.ToggleSelect()` is called
3. `SelectionViewModel.SelectBlock()` updates selection state
4. Selection listeners are notified
5. All block views update their visual state
6. `ManagerWindowViewModel` updates selected save game title

### Copying a Block

1. User clicks "Copy" button
2. `ManagerWindowViewModel.CopyCommand()` is called
3. `MemoryCard.CopyBlockTo()` copies block data (TODO: implementation)
4. `MemoryCard.Write()` saves changes to file
5. UI bindings are refreshed

### Deleting a Block

1. User clicks "Delete" button
2. `ManagerWindowViewModel.DeleteCommand()` is called
3. `MemoryCard.DeleteBlockFrom()` marks block as deleted
4. `MemoryCard.Write()` saves changes to file
5. `RefreshCardBindings()` updates UI

## Key Design Decisions

### 1. MVVM Pattern

- **Rationale**: Separates UI concerns from business logic
- **Benefits**: 
  - Testable view models
  - Reusable domain logic
  - Clear data flow

### 2. Data Binding

- **Rationale**: Fyne's binding system enables reactive UI updates
- **Benefits**:
  - Automatic UI synchronization
  - Reduced manual refresh calls
  - Type-safe bindings

### 3. Dependency Injection

- **Rationale**: Centralized dependency management
- **Benefits**:
  - Loose coupling
  - Easier testing
  - Clear dependency graph

### 4. Domain Layer Separation

- **Rationale**: PSX memory card format logic is independent of UI
- **Benefits**:
  - Reusable in other contexts
  - Easier to test
  - Clear format specification

## File Format

The application works with PSX memory card files (.mcr format):

- **Total Size**: 128 KB (131,072 bytes)
- **Block Size**: 8 KB (8,192 bytes)
- **Number of Blocks**: 15 data blocks (16th block is header)
- **Block Structure**:
  - Title frame (128 bytes)
  - Icon frames (3 x 128 bytes)
  - Data frames (60 x 128 bytes)

## Dependencies

### External Libraries

- **Fyne** (`fyne.io/fyne/v2`): Cross-platform GUI framework
- **uber/dig** (`go.uber.org/dig`): Dependency injection
- **golang.org/x/text**: Text encoding (Shift-JIS support)

## Testing Strategy

The codebase follows a test-driven development approach:

- Unit tests for individual functions/methods
- Tests for domain logic (memory card format)
- Avoid mocking filesystem (use temp directories)
- Avoid mocking loggers
- One test block per method

## Known Limitations

1. **Block Copying**: `CopyBlockTo()` method is not fully implemented (TODO)
2. **Error Handling**: Some error cases may need more robust handling
3. **Thread Safety**: Selection model has a TODO comment about atomic operations

## Future Enhancements

- Complete block copying implementation
- Support for drag-and-drop file loading
- Export/import individual save games
- Memory card validation and repair
- Support for multiple memory card formats
