package memcard

import (
	"encoding/binary"
	"fmt"
	"os"
)

func (mc *MemoryCard) Write(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, mc)
	if err != nil {
		return fmt.Errorf("failed to write memory card to file: %w", err)
	}

	return nil
}
