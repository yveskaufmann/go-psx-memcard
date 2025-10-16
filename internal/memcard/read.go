package memcard

import (
	"encoding/binary"
	"errors"
	"os"
)

var (
	ErrInvalidMemoryCardSize = errors.New("invalid memory card size, expected 128 Kilobytes")
	ErrEmptyFile             = errors.New("file is empty")
)

func Open(filePath string) (*MemoryCard, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.Size() == 0 {
		return nil, ErrEmptyFile
	}

	if fileInfo.Size() != MemoryCardTotalSize {
		return nil, ErrInvalidMemoryCardSize
	}

	var memCard MemoryCard
	err = binary.Read(file, binary.LittleEndian, &memCard)
	if err != nil {
		return nil, err
	}

	return &memCard, nil
}
