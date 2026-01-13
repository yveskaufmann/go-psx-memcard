package memcard

import (
	"errors"
)

var (
	ErrInvalidBlockIndex    = errors.New("invalid block index")
	ErrSourceBlockNotInUse  = errors.New("source block is not in use")
	ErrNoFreeBlockAvailable = errors.New("no free block available on target memory card")
	ErrTargetCardNil        = errors.New("target memory card is nil")
)

// FindFreeBlock finds the first available (free) block on the memory card.
// Returns the block index (0-14) and true if a free block was found, or -1 and false if no free block is available.
func (mc *MemoryCard) FindFreeBlock() (int, bool) {
	for i := 0; i < NumBlocks; i++ {
		state := mc.DirectoryFrames[i].BlockAllocationState
		if state == BlockAllocationStateFreeFresh ||
			state == BlockAllocationStateFreeDeletedFirst ||
			state == BlockAllocationStateFreeDeletedMiddle ||
			state == BlockAllocationStateFreeDeletedLast {
			return i, true
		}
	}
	return -1, false
}

// CountBlocks returns the total, used, and free block counts for the memory card.
func (mc *MemoryCard) CountBlocks() (total, used, free int) {
	total = NumBlocks
	used = 0
	free = 0

	for i := 0; i < NumBlocks; i++ {
		state := mc.DirectoryFrames[i].BlockAllocationState
		if state == BlockAllocationStateFreeFresh ||
			state == BlockAllocationStateFreeDeletedFirst ||
			state == BlockAllocationStateFreeDeletedMiddle ||
			state == BlockAllocationStateFreeDeletedLast {
			free++
		} else {
			used++
		}
	}

	return total, used, free
}

// CopyBlockTo copies a block from the source memory card to the target memory card.
// It finds a free block on the target card, copies the block data and directory frame,
// and updates the allocation state to indicate it's a first-or-only block (0x51).
func (mc *MemoryCard) CopyBlockTo(blockIndex int, targetCard *MemoryCard) error {
	if targetCard == nil {
		return ErrTargetCardNil
	}

	if blockIndex < 0 || blockIndex >= NumBlocks {
		return ErrInvalidBlockIndex
	}

	// Verify source block is in use
	sourceDirFrame := mc.DirectoryFrames[blockIndex]
	if sourceDirFrame.BlockAllocationState == BlockAllocationStateFreeFresh ||
		sourceDirFrame.BlockAllocationState == BlockAllocationStateFreeDeletedFirst ||
		sourceDirFrame.BlockAllocationState == BlockAllocationStateFreeDeletedMiddle ||
		sourceDirFrame.BlockAllocationState == BlockAllocationStateFreeDeletedLast {
		return ErrSourceBlockNotInUse
	}

	// Find a free block on the target card
	targetBlockIndex, found := targetCard.FindFreeBlock()
	if !found {
		return ErrNoFreeBlockAvailable
	}

	// Copy the block data
	targetCard.Blocks[targetBlockIndex] = mc.Blocks[blockIndex]

	// Copy the directory frame, but update it for the target
	targetCard.DirectoryFrames[targetBlockIndex] = sourceDirFrame

	// Update the allocation state to indicate it's a first-or-only block
	// When copying a single block, it becomes a standalone file
	targetCard.DirectoryFrames[targetBlockIndex].BlockAllocationState = BlockAllocationStateInUseFirstOnlyBlock

	// Update NextBlock to FFFFh (indicates last-or-only block)
	targetCard.DirectoryFrames[targetBlockIndex].NextBlock = 0xFFFF

	// Update file size to 0x2000 (8KB) for a single block file
	targetCard.DirectoryFrames[targetBlockIndex].FileSize = BlockSize

	// Update the block number in the title frame to match the new block index
	targetCard.Blocks[targetBlockIndex].TitleFrame.BlockNumber = byte(targetBlockIndex + 1)

	// Recalculate and update the checksum for the directory frame
	targetCard.DirectoryFrames[targetBlockIndex].Checksum = calculateDirectoryFrameChecksum(&targetCard.DirectoryFrames[targetBlockIndex])

	return nil
}

func (mc *MemoryCard) DeleteBlockFrom(blockIndex int) error {
	// TODO: Implement block deletion logic

	mc.DirectoryFrames[blockIndex].BlockAllocationState = BlockAllocationStateFreeDeletedFirst
	mc.DirectoryFrames[blockIndex].FileName = NewEmptyFileName()

	mc.Blocks[blockIndex].CleanBlock()

	return nil
}
