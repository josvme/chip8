package utils

// GetNibbles This takes byte of length 2
func GetNibbles(inst []byte) (byte, byte, byte, byte) {
	if len(inst) != 2 {
		panic("Pass 2 byte instructions only")
	}
	b1 := inst[0]
	b2 := inst[1]

	a := (b1 & 0b11110000) >> 4
	b := b1 & 0b00001111
	c := (b2 & 0b11110000) >> 4
	d := b2 & 0b00001111

	return a, b, c, d
}

func ModFn(x int, n int) int {
	return ((x % n) + n) % n
}
