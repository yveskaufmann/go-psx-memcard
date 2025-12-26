package memcard

func (mc *MemoryCard) CopyBlockTo(blockIndex int, targetCard *MemoryCard) error {
	// TODO: Implement block copying logic
	return nil
}

func (mc *MemoryCard) DeleteBlockFrom(blockIndex int) error {
	// TODO: Implement block deletion logic

	mc.DirectoryFrames[blockIndex].BlockAllocationState = BlockAllocationStateFreeDeletedFirst
	mc.DirectoryFrames[blockIndex].FileName = NewEmptyFileName()

	mc.Blocks[blockIndex].CleanBlock()

	return nil
}
