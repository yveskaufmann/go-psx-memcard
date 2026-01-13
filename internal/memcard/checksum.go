package memcard

// Checksum provides centralized checksum calculation functions for PSX memory card format.
// All PSX memory card frames use XOR checksum algorithm where the checksum byte
// is calculated by XORing all other bytes in the frame.

// calculateXORChecksum calculates the XOR checksum for a byte slice.
// This is the core algorithm used by all PSX memory card frame checksums.
func calculateXORChecksum(data []byte) byte {
	var checksum byte = 0
	for i := 0; i < len(data); i++ {
		checksum ^= data[i]
	}
	return checksum
}

// calculateHeaderChecksum calculates the XOR checksum for the header frame.
// The checksum is calculated by XORing all bytes except the checksum byte itself (at offset 0x7F).
func calculateHeaderChecksum(header *HeaderFrame) byte {
	var checksum byte = 0
	checksum ^= header.MagicBytes[0]
	checksum ^= header.MagicBytes[1]
	for i := 0; i < len(header.Unused); i++ {
		checksum ^= header.Unused[i]
	}
	return checksum
}

// calculateDirectoryFrameChecksum calculates the XOR checksum for a directory frame.
// The checksum is calculated by XORing all bytes in the frame except the checksum byte itself (at offset 0x7F).
func calculateDirectoryFrameChecksum(frame *DirectoryFrame) byte {
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

// calculateBrokenSelectorChecksum calculates the XOR checksum for a broken selector frame.
func calculateBrokenSelectorChecksum(selector *BrokenSelector) byte {
	var checksum byte = 0
	// XOR BrokenSectorNumber (4 bytes)
	checksum ^= byte(selector.BrokenSectorNumber)
	checksum ^= byte(selector.BrokenSectorNumber >> 8)
	checksum ^= byte(selector.BrokenSectorNumber >> 16)
	checksum ^= byte(selector.BrokenSectorNumber >> 24)
	// XOR Reserved bytes
	for i := 0; i < len(selector.Reserved); i++ {
		checksum ^= selector.Reserved[i]
	}
	return checksum
}

// calculateFrameChecksum calculates the XOR checksum for a 128-byte frame.
// It XORs all bytes in the provided slice.
func calculateFrameChecksum(frame []byte) byte {
	return calculateXORChecksum(frame)
}
