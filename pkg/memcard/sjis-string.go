package memcard

import (
	"bytes"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type ShiftJISString struct {
	Data [64]byte
}

func NewShiftJISString(str string) (ShiftJISString, error) {
	encoder := japanese.ShiftJIS.NewEncoder()
	encodedStr, _, err := transform.String(encoder, str)
	if err != nil {
		return ShiftJISString{}, err
	}

	var data [64]byte
	copy(data[:], encodedStr)

	return ShiftJISString{Data: data}, nil
}

func (s *ShiftJISString) String() string {

	decoder := japanese.ShiftJIS.NewDecoder()

	nullByteIndex := bytes.Index(s.Data[:], []byte{0})
	if nullByteIndex == -1 {
		nullByteIndex = len(s.Data)
	}

	str, _, err := transform.String(decoder, string(s.Data[:nullByteIndex]))
	if err != nil {
		return ""
	}

	return strings.TrimSpace(str)
}
