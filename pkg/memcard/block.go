package memcard

import (
	"errors"
	"image"
)

var (
	ErrInvalidBlockNumber = errors.New("invalid block number")
)

type BlockItem struct {
	Title       string
	Icon        image.Image
	BlockNumber uint8
}

func (mc *MemoryCard) GetBlock(blockNumber int) (*BlockItem, error) {
	if blockNumber < 0 || blockNumber >= NumBlocks {
		return nil, ErrInvalidBlockNumber
	}

	df := mc.DirectoryFrames[blockNumber]
	block := mc.Blocks[blockNumber]

	if df.BlockAllocationState == BlockAllocationStateFreeFresh ||
		df.BlockAllocationState == BlockAllocationStateFreeDeletedFirst ||
		df.BlockAllocationState == BlockAllocationStateFreeDeletedMiddle ||
		df.BlockAllocationState == BlockAllocationStateFreeDeletedLast {
		return nil, nil // Block is free or deleted
	}

	item := &BlockItem{
		Title:       block.TitleFrame.Title.String(),
		Icon:        block.IconFrames[0].ToImage(block.TitleFrame.IconColorPalette),
		BlockNumber: uint8(blockNumber),
	}

	return item, nil
}

func (mc *MemoryCard) ListBlocks() ([]BlockItem, error) {
	var items []BlockItem

	for i := range NumBlocks {
		item, err := mc.GetBlock(i)
		if err != nil {
			return nil, err
		}
		if item != nil {
			items = append(items, *item)
		}
	}

	return items, nil
}
