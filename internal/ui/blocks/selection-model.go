package blocks

import (
	"sync"

	"com.yv35.memcard/internal/memcard"
	"fyne.io/fyne/v2/data/binding"
)

const NoBlockSelected = -1
const NoCardSelected = ""

type SelectionListener interface {
	SelectionChanged(cardID memcard.MemoryCardID, blockIndex int)
}

type funcSelectionListener struct {
	onSelectionChanged func(cardID memcard.MemoryCardID, blockIndex int)
}

func NewSelectionChangedListener(onSelectionChanged func(cardID memcard.MemoryCardID, blockIndex int)) SelectionListener {
	return &funcSelectionListener{
		onSelectionChanged: onSelectionChanged,
	}
}

func (s *funcSelectionListener) SelectionChanged(cardID memcard.MemoryCardID, blockIndex int) {
	s.onSelectionChanged(cardID, blockIndex)
}

type SelectionViewModel struct {
	SelectedBlockIndex binding.Int
	SelectedCardId     binding.String
	listeners          []SelectionListener
	lock               sync.RWMutex
}

func NewBlockSelectionViewModel() *SelectionViewModel {
	selection := &SelectionViewModel{
		SelectedBlockIndex: binding.NewInt(),
		SelectedCardId:     binding.NewString(),
	}

	selection.SelectedBlockIndex.Set(NoBlockSelected)
	selection.SelectedCardId.Set(string(NoCardSelected))

	return selection
}

func (b *SelectionViewModel) ClearSelection() {
	b.SelectedBlockIndex.Set(NoBlockSelected)
	b.SelectedCardId.Set(string(NoCardSelected))

	b.notifySelectionChanged()
}

func (b *SelectionViewModel) SelectBlock(cardID memcard.MemoryCardID, blockIndex int) {
	if cardID == NoCardSelected || blockIndex == NoBlockSelected {
		b.ClearSelection()
		return
	}

	// TODO: This must be atomic otherwise listeners may see inconsistent state
	b.SelectedCardId.Set(string(cardID))
	b.SelectedBlockIndex.Set(blockIndex)

	b.notifySelectionChanged()
}

func (b *SelectionViewModel) UnselectBlock(cardID memcard.MemoryCardID, blockIndex int) {
	b.ClearSelection()
}

func (b *SelectionViewModel) BlockIndex() int {
	blockIndex, err := b.SelectedBlockIndex.Get()
	if err != nil || blockIndex == NoBlockSelected {
		return NoBlockSelected
	}

	return blockIndex
}

func (b *SelectionViewModel) CardId() memcard.MemoryCardID {
	cardIDStr, err := b.SelectedCardId.Get()
	if err != nil || cardIDStr == NoCardSelected {
		return memcard.MemoryCardID(NoCardSelected)
	}
	return memcard.MemoryCardID(cardIDStr)
}

func (b *SelectionViewModel) AddListener(listener SelectionListener) {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.listeners = append(b.listeners, listener)
}

func (b *SelectionViewModel) RemoveListener(listenerToRemove SelectionListener) {
	b.lock.Lock()
	defer b.lock.Unlock()

	var newListenerList []SelectionListener
	for i := range b.listeners {
		listener := b.listeners[i]
		if listener == listenerToRemove {
			continue
		}
		newListenerList = append(newListenerList, listener)
	}

	b.listeners = newListenerList
}

func (b *SelectionViewModel) notifySelectionChanged() {
	b.lock.RLock()
	defer b.lock.RUnlock()

	cardID := b.CardId()
	blockIndex := b.BlockIndex()

	for i := range b.listeners {
		listener := b.listeners[i]
		listener.SelectionChanged(cardID, blockIndex)
	}
}
