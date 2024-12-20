package core

import (
	"encoding/hex"
	"fmt"
)

const registerVF = 15
const programStart = 512

type Cpu struct {
	internals CpuInternals
	display   IDisplay
	sound     ISound
	memory    *Memory
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

func NewCpu(display IDisplay, sound ISound, memory *Memory) Cpu {
	return Cpu{
		internals: CpuInternals{
			V:     [16]byte{},
			I:     [2]byte{},
			Delay: 0,
			Sound: 0,
			PC:    programStart,
			SP:    0,
			Stack: [16]uint16{},
		},
		display: display,
		sound:   sound,
		memory:  memory,
	}
}

func (cpu *Cpu) Run() {
	for {
		pc := cpu.internals.PC
		inst := cpu.memory.Mem[pc : pc+2]
		fmt.Println(pc, hex.EncodeToString(inst))
		cpu.executeInstruction(inst)
	}
}

func (cpu *Cpu) executeInstruction(inst []byte) {
	switch {
	case hex.EncodeToString(inst) == "00e0":
		cpu.display.clearScreen()
		cpu.internals.PC += 2
	}
	//time.Sleep(10 * time.Millisecond)
}
