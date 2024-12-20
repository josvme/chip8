package utils_test

import (
	"chip8/utils"
	"testing"
)

func TestGetNibbles(t *testing.T) {
	v := []byte{0x11, 0x0A}
	a, b, c, d := utils.GetNibbles(v)
	print(a, b, c, d)
	if a != 0x01 || b != 0x01 || c != 0x00 || d != 0x0A {
		t.Errorf("Error in Nibble")
	}
}
