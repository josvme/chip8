package core

import (
	"chip8/utils"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

const programStart = 512
const F = 15

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
		time.Sleep(10 * time.Millisecond)
	}
}

func (cpu *Cpu) executeInstruction(inst []byte) {
	state := &cpu.internals
	high, x, y, low := utils.GetNibbles(inst)
	hexString := hex.EncodeToString(inst)
	switch {
	case hexString == "00e0":
		// Clear Screen
		cpu.display.clearScreen()
		state.PC += 2
	case hexString == "00ee":
		// Return
		state.PC = state.Stack[cpu.internals.SP-1]
		state.SP -= 1
	case high == 1:
		state.PC = binary.BigEndian.Uint16([]byte{x, inst[1]})
	case high == 2:
		state.Stack[state.SP] = state.PC
		state.SP += 1
		state.PC = binary.BigEndian.Uint16([]byte{x, inst[1]})
	case high == 3:
		if state.V[x] == inst[1] {
			state.PC += 2
		}
		state.PC += 2
	case high == 4:
		if state.V[x] != inst[1] {
			state.PC += 2
		}
		state.PC += 2
	case high == 5:
		if state.V[x] == state.V[y] {
			state.PC += 2
		}
		state.PC += 2
	case high == 6:
		state.V[x] = inst[1]
		state.PC += 2
	case high == 7:
		state.V[x] = state.V[x] + inst[1]
		state.PC += 2
	case high == 8 && low == 0:
		state.V[x] = state.V[y]
		state.PC += 2
	case high == 8 && low == 1:
		state.V[x] = state.V[x] | state.V[y]
		state.PC += 2
	case high == 8 && low == 2:
		state.V[x] = state.V[x] & state.V[y]
		state.PC += 2
	case high == 8 && low == 3:
		state.V[x] = state.V[x] ^ state.V[y]
		state.PC += 2
	case high == 8 && low == 4:
		state.V[x] = state.V[x] + state.V[y]
		state.PC += 2
	case high == 8 && low == 5:
		state.V[x] = state.V[x] - state.V[y]
		state.PC += 2
	case high == 8 && low == 6:
		state.V[F] = state.V[x] & 0b00000001
		state.V[x] = state.V[x] / 2
		state.PC += 2
	case high == 8 && low == 7:
		if state.V[y] > state.V[x] {
			state.V[F] = 1
		} else {
			state.V[F] = 0
		}
		state.V[x] = state.V[y] - state.V[x]
		state.PC += 2
	case high == 8 && low == 0xE:
		state.V[F] = state.V[x] & 0b10000000
		state.V[x] = state.V[x] * 2
		state.PC += 2
	case high == 9 && low == 0:
		if state.V[x] != state.V[y] {
			state.PC += 2
		}
		state.PC += 2
	}
}
