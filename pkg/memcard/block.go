package memcard

import (
	"errors"
	"image"

	animatedsprite "com.yvka.memcard/pkg/ui/animated-sprite"
)

var (
	ErrInvalidBlockNumber = errors.New("invalid block number")
)

type BlockItem struct {
	Title       string
	Animation   animatedsprite.Animation
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

	frames := []image.Image{}
	for idx, f := range block.IconFrames {

		switch block.TitleFrame.IconDisplayFlag {
		case IconDisplayFlagOneFrameIcon:
			if idx > 0 {
				continue
			}
		case IconDisplayFlagTwoFrameIcon:
			if idx > 1 {
				continue
			}
		case IconDisplayFlagThreeFrameIcon:
			if idx > 2 {
				continue
			}

		}
		frames = append(frames, f.ToImage(block.TitleFrame.IconColorPalette))
	}

	item := &BlockItem{
		Title:     block.TitleFrame.Title.String(),
		Animation: animatedsprite.NewAnimation(frames),

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
