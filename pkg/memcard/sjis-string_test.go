package memcard

import "testing"

func TestShiftJISString_String(t *testing.T) {
	shiftJISString := ShiftJISString{
		Data: [64]byte{
			130, 96, 130, 98, 130, 100, 130, 98, 130, 110, 130, 108, 130, 97, 130, 96, 130, 115, 130, 82, 129, 64, 130, 133, 130, 140, 130, 133, 130, 131, 130, 148, 130, 146, 130, 143, 130, 147, 130, 144, 130, 136, 130, 133, 130, 146, 130, 133, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}

	expected := "ＡＣＥＣＯＭＢＡＴ３　ｅｌｅｃｔｒｏｓｐｈｅｒｅ"
	result := shiftJISString.String()

	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}

func TestNewShiftJISString(t *testing.T) {
	original := "ＡＣＥＣＯＭＢＡＴ３　ｅｌｅｃｔｒｏｓｐｈｅｒｅ"
	sjisString, err := NewShiftJISString(original)
	if err != nil {
		t.Fatalf("Error encoding string: %v", err)
	}

	result := sjisString.String()
	if result != original {
		t.Errorf("Expected: %s, but got: %s", original, result)
	}
}
