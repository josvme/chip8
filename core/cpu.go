package core

import "encoding/hex"

const registerVF = 15

type Cpu struct {
	internals CpuInternals
	display   IDisplay
	sound     ISound
}
type CpuInternals struct {
	V     [16]byte
	I     [2]byte
	Delay byte
	Sound byte
	PC    uint16
	SP    uint8
	Stack [16]uint16
}

func NewCpu() Cpu {
	return Cpu{}
}

func (cpu *Cpu) executeInstruction(inst [2]byte) {
	switch {
	case hex.EncodeToString(inst[:]) == "00e0":
		cpu.display.clearScreen()
	}
}
