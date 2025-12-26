package memcard

type MemoryCardID string

const (
	MemoryCardLeft  MemoryCardID = "Card-1"
	MemoryCardRight MemoryCardID = "Card-2"
)

type IconDisplayFlag byte

const (
	IconDisplayFlagOneFrameIcon   IconDisplayFlag = 0x11
	IconDisplayFlagTwoFrameIcon   IconDisplayFlag = 0x12
	IconDisplayFlagThreeFrameIcon IconDisplayFlag = 0x13
)

type BlockAllocationState uint32

const (
	BlockAllocationStateInUseFirstOnlyBlock BlockAllocationState = 0x51 // first-or-only block of a file
	BlockAllocationStateInUseMiddleBlock    BlockAllocationState = 0x52 // middle block of a file (if 3 or more blocks)
	BlockAllocationStateInUseLastBlock      BlockAllocationState = 0x53 // last block of a file (if 2 or more blocks)
	BlockAllocationStateFreeFresh           BlockAllocationState = 0xA0 // freshly formatted
	BlockAllocationStateFreeDeletedFirst    BlockAllocationState = 0xA1 // deleted (first-or-only block of file)
	BlockAllocationStateFreeDeletedMiddle   BlockAllocationState = 0xA2 // deleted (middle block of file)
	BlockAllocationStateFreeDeletedLast     BlockAllocationState = 0xA3 // deleted (last block of file)
)

const (
	MemoryCardTotalSize = 131072 // 128 KB
	BlockSize           = 8192   // 8 KB
	NumBlocks           = 15     // 15 blocks (the 16th block is used for the header
	FrameSize           = 128    // 128 bytes per frame
)

type MemoryCard struct {
	Header                     HeaderFrame
	DirectoryFrames            [15]DirectoryFrame
	BrokenSelectors            [20]BrokenSelector
	BrokenSelectorReplacements [20]BrokenSelectorReplacement
	UnusedFrames               [7][128]byte
	WriteTestFrame             [128]byte
	Blocks                     [15]Block
}

type HeaderFrame struct {
	MagicBytes [2]byte
	Unused     [125]byte
	Checksum   byte
}

type FileName [21]byte

func NewEmptyFileName() FileName {
	// TODO: Check what are the correct bytes for an empty filename
	var fn FileName
	for i := range fn {
		fn[i] = 0x0 // fill with null bytes
	}
	return fn
}

func (f *FileName) Region() string {
	regionCode := string(f[0:2])

	switch regionCode {
	case "BI":
		return "Japan"
	case "BE":
		return "Europe"
	case "BA":
		return "America"
	}
	return "Unknown"
}

func (f *FileName) GameCode() string {
	gameCode := string(f[2:12])
	return gameCode
}

func (f *FileName) GameName() string {
	gameName := string(f[12:])
	return gameName
}

type DirectoryFrame struct {
	BlockAllocationState BlockAllocationState
	FileSize             uint32
	NextBlock            uint16
	FileName             FileName
	Zero                 byte
	Reserved             [95]byte
	Checksum             byte
}

type BrokenSelector struct {
	BrokenSectorNumber uint32
	Reserved           [123]byte
	Checksum           byte
}

type BrokenSelectorReplacement struct {
	Date [128]byte
}

type Block struct {
	TitleFrame BlockTitleFrame
	IconFrames [3]IconBitmapFrame
	Data       [60]DataFrame
}

func (b *Block) CleanBlock() {
	var emptyTitleFrame BlockTitleFrame
	var emptyIconFrame IconBitmapFrame
	var emptyDataFrame DataFrame

	b.TitleFrame = emptyTitleFrame
	for i := range b.IconFrames {
		b.IconFrames[i] = emptyIconFrame
	}
	for i := range b.Data {
		b.Data[i] = emptyDataFrame
	}
}

type BlockTitleFrame struct {
	ID               [2]byte
	IconDisplayFlag  IconDisplayFlag
	BlockNumber      byte
	Title            ShiftJISString
	Reserved         [28]byte
	IconColorPalette [16]uint16
}

type DataFrame [128]byte
