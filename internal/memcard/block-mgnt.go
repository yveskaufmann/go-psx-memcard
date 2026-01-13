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

// calculateDirectoryFrameChecksum calculates the XOR checksum for a directory frame.
// The checksum is calculated by XORing all bytes in the frame except the checksum byte itself (at offset 0x7F).
func calculateDirectoryFrameChecksum(frame *DirectoryFrame) byte {
	// Convert frame to byte array for checksum calculation
	// We need to XOR all bytes except the checksum byte at offset 0x7F
	var checksum byte = 0

	// XOR BlockAllocationState (4 bytes: 0x00-0x03)
	checksum ^= byte(frame.BlockAllocationState)
	checksum ^= byte(frame.BlockAllocationState >> 8)
	checksum ^= byte(frame.BlockAllocationState >> 16)
	checksum ^= byte(frame.BlockAllocationState >> 24)

	// XOR FileSize (4 bytes: 0x04-0x07)
	checksum ^= byte(frame.FileSize)
	checksum ^= byte(frame.FileSize >> 8)
	checksum ^= byte(frame.FileSize >> 16)
	checksum ^= byte(frame.FileSize >> 24)

	// XOR NextBlock (2 bytes: 0x08-0x09)
	checksum ^= byte(frame.NextBlock)
	checksum ^= byte(frame.NextBlock >> 8)

	// XOR FileName (21 bytes: 0x0A-0x1E)
	for i := 0; i < len(frame.FileName); i++ {
		checksum ^= frame.FileName[i]
	}

	// XOR Zero byte (0x1F)
	checksum ^= frame.Zero

	// XOR Reserved (95 bytes: 0x20-0x7E)
	for i := 0; i < len(frame.Reserved); i++ {
		checksum ^= frame.Reserved[i]
	}

	return checksum
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
