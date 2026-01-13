package memcard

// NewFormattedMemoryCard creates a new, properly formatted empty memory card.
// The card is initialized with:
// - Header frame with "MC" magic bytes and correct checksum (0x0E)
// - All directory frames set to free/freshly formatted state (0xA0)
// - All blocks empty
// - Broken selectors set to 0xFFFFFFFF (no broken sectors)
// - Unused frames filled with 0xFF
// - Write test frame initialized
func NewFormattedMemoryCard() *MemoryCard {
	card := &MemoryCard{}

	// Initialize header frame
	card.Header.MagicBytes[0] = 'M'
	card.Header.MagicBytes[1] = 'C'
	// Unused bytes are already zero-initialized
	// Calculate header checksum: XOR of all bytes except checksum byte
	card.Header.Checksum = calculateHeaderChecksum(&card.Header)

	// Initialize all directory frames as free/freshly formatted
	for i := 0; i < NumBlocks; i++ {
		card.DirectoryFrames[i].BlockAllocationState = BlockAllocationStateFreeFresh
		card.DirectoryFrames[i].FileSize = 0
		card.DirectoryFrames[i].NextBlock = 0xFFFF
		card.DirectoryFrames[i].FileName = NewEmptyFileName()
		card.DirectoryFrames[i].Zero = 0
		// Reserved bytes are already zero-initialized
		// Calculate directory frame checksum
		card.DirectoryFrames[i].Checksum = calculateDirectoryFrameChecksum(&card.DirectoryFrames[i])
	}

	// Initialize broken selectors (all set to 0xFFFFFFFF = no broken sectors)
	for i := 0; i < len(card.BrokenSelectors); i++ {
		card.BrokenSelectors[i].BrokenSectorNumber = 0xFFFFFFFF
		// Reserved bytes are already zero-initialized
		// Calculate broken selector checksum
		card.BrokenSelectors[i].Checksum = calculateBrokenSelectorChecksum(&card.BrokenSelectors[i])
	}

	// Initialize broken selector replacements (filled with 0xFF)
	for i := 0; i < len(card.BrokenSelectorReplacements); i++ {
		for j := 0; j < len(card.BrokenSelectorReplacements[i].Date); j++ {
			card.BrokenSelectorReplacements[i].Date[j] = 0xFF
		}
	}

	// Initialize unused frames (filled with 0xFF)
	for i := 0; i < len(card.UnusedFrames); i++ {
		for j := 0; j < len(card.UnusedFrames[i]); j++ {
			card.UnusedFrames[i][j] = 0xFF
		}
	}

	// Initialize write test frame (same as block 0 header)
	card.WriteTestFrame[0] = 'M'
	card.WriteTestFrame[1] = 'C'
	for i := 2; i < len(card.WriteTestFrame)-1; i++ {
		card.WriteTestFrame[i] = 0
	}
	// Calculate checksum for write test frame
	card.WriteTestFrame[127] = calculateFrameChecksum(card.WriteTestFrame[:127])

	// Blocks are already zero-initialized (empty)

	return card
}
