package blockstats

import (
	"fyne.io/fyne/v2/data/binding"
)

// MemoryCardBlockStatsViewModel manages the state and logic for displaying memory card block statistics.
type MemoryCardBlockStatsViewModel struct {
	leftCardTotal  binding.Int
	leftCardUsed   binding.Int
	leftCardFree   binding.Int
	rightCardTotal binding.Int
	rightCardUsed  binding.Int
	rightCardFree  binding.Int

	getLeftCardStats  func() (total, used, free int)
	getRightCardStats func() (total, used, free int)
}

// NewMemoryCardBlockStatsViewModel creates a new view model for memory card block statistics.
func NewMemoryCardBlockStatsViewModel(
	getLeftCardStats func() (total, used, free int),
	getRightCardStats func() (total, used, free int),
) *MemoryCardBlockStatsViewModel {
	vm := &MemoryCardBlockStatsViewModel{
		leftCardTotal:     binding.NewInt(),
		leftCardUsed:      binding.NewInt(),
		leftCardFree:      binding.NewInt(),
		rightCardTotal:    binding.NewInt(),
		rightCardUsed:     binding.NewInt(),
		rightCardFree:     binding.NewInt(),
		getLeftCardStats:  getLeftCardStats,
		getRightCardStats: getRightCardStats,
	}

	// Initialize with zero values
	vm.leftCardTotal.Set(0)
	vm.leftCardUsed.Set(0)
	vm.leftCardFree.Set(0)
	vm.rightCardTotal.Set(0)
	vm.rightCardUsed.Set(0)
	vm.rightCardFree.Set(0)

	return vm
}

// UpdateStatistics refreshes the statistics from the memory cards.
func (vm *MemoryCardBlockStatsViewModel) UpdateStatistics() {
	leftTotal, leftUsed, leftFree := vm.getLeftCardStats()
	rightTotal, rightUsed, rightFree := vm.getRightCardStats()

	vm.leftCardTotal.Set(leftTotal)
	vm.leftCardUsed.Set(leftUsed)
	vm.leftCardFree.Set(leftFree)
	vm.rightCardTotal.Set(rightTotal)
	vm.rightCardUsed.Set(rightUsed)
	vm.rightCardFree.Set(rightFree)
}

// LeftCardTotal returns the binding for left card total blocks.
func (vm *MemoryCardBlockStatsViewModel) LeftCardTotal() binding.Int {
	return vm.leftCardTotal
}

// LeftCardUsed returns the binding for left card used blocks.
func (vm *MemoryCardBlockStatsViewModel) LeftCardUsed() binding.Int {
	return vm.leftCardUsed
}

// LeftCardFree returns the binding for left card free blocks.
func (vm *MemoryCardBlockStatsViewModel) LeftCardFree() binding.Int {
	return vm.leftCardFree
}

// RightCardTotal returns the binding for right card total blocks.
func (vm *MemoryCardBlockStatsViewModel) RightCardTotal() binding.Int {
	return vm.rightCardTotal
}

// RightCardUsed returns the binding for right card used blocks.
func (vm *MemoryCardBlockStatsViewModel) RightCardUsed() binding.Int {
	return vm.rightCardUsed
}

// RightCardFree returns the binding for right card free blocks.
func (vm *MemoryCardBlockStatsViewModel) RightCardFree() binding.Int {
	return vm.rightCardFree
}
