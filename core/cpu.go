package core

import (
	"chip8/utils"
	"crypto/rand"
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
	input     IInput
	keyMap    map[byte]uint16
}
type CpuInternals struct {
	V     [16]byte
	I     uint16
	Delay byte
	Sound byte
	PC    uint16
	SP    uint8
	Stack [16]uint16
}

func NewCpu(display IDisplay, sound ISound, memory *Memory) Cpu {
	keymap := make(map[byte]uint16)
	for i := 0; i < 16; i++ {
		keymap[byte(i)] = uint16(i * 5)
	}
	return Cpu{
		internals: CpuInternals{
			V:     [16]byte{},
			I:     0,
			Delay: 0,
			Sound: 0,
			PC:    programStart,
			SP:    0,
			Stack: [16]uint16{},
		},
		display: display,
		sound:   sound,
		memory:  memory,
		keyMap:  keymap,
	}
}

func (cpu *Cpu) Run() {
	for {
		pc := cpu.internals.PC
		inst := cpu.memory.Mem[pc : pc+2]
		fmt.Println(pc, hex.EncodeToString(inst))
		cpu.executeInstruction(inst)
		time.Sleep(2 * time.Millisecond)
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
		// Update flags
		if state.V[x] > state.V[y] {
			state.V[F] = 1
		} else {
			state.V[F] = 0
		}
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
		msb := state.V[x] & 0b10000000
		// We only need 1 bit as it is a flag and hence need to shift
		state.V[F] = msb >> 7
		state.V[x] = state.V[x] << 1
		state.PC += 2
	case high == 9 && low == 0:
		if state.V[x] != state.V[y] {
			state.PC += 2
		}
		state.PC += 2
	case high == 0xA:
		state.I = binary.BigEndian.Uint16([]byte{x, inst[1]})
		state.PC += 2
	case high == 0xB:
		state.PC = binary.BigEndian.Uint16([]byte{x, inst[1]}) + uint16(state.V[0])
	case high == 0xC:
		r := make([]byte, 1)
		rand.Read(r)
		state.V[x] = inst[1] & r[0]
		state.PC += 2
	case high == 0xD:
		vF := uint8(0)
		for i := uint8(0); i < low; i++ {
			vF = vF | cpu.display.drawPixels(state.V[x], state.V[y]+i, cpu.memory.Mem[state.I+uint16(i)])
		}
		// Set collision flag
		state.V[F] = vF
		state.PC += 2
	case high == 0xE && inst[1] == 0x9E:
		if cpu.input.isKeyPressed(state.V[x]) {
			state.PC += 2
		}
		state.PC += 2
	case high == 0xE && inst[1] == 0xA1:
		if !cpu.input.isKeyPressed(state.V[x]) {
			state.PC += 2
		}
		state.PC += 2
	case high == 0xF && inst[1] == 0x07:
		state.V[x] = state.Delay
		state.PC += 2
	case high == 0xF && inst[1] == 0x0A:
		scanKey := cpu.input.scanKey()
		state.V[x] = scanKey
		state.PC += 2
	case high == 0xF && inst[1] == 0x15:
		state.Delay = x
		state.PC += 2
	case high == 0xF && inst[1] == 0x18:
		state.Sound = state.V[x]
		state.PC += 2
	case high == 0xF && inst[1] == 0x1E:
		state.I = state.I + uint16(state.V[x])
		state.PC += 2
	case high == 0xF && inst[1] == 0x29:
		state.I = cpu.keyMap[state.V[x]]
		state.PC += 2
	case high == 0xF && inst[1] == 0x33:
		// Already in uint8
		bcd := state.V[x]
		cpu.memory.Mem[state.I] = bcd / 100
		cpu.memory.Mem[state.I+1] = (bcd / 10) % 10
		cpu.memory.Mem[state.I+2] = bcd % 10
		state.PC += 2
	case high == 0xF && inst[1] == 0x55:
		for i := 0; i <= int(x); i++ {
			cpu.memory.Mem[int(state.I)+i] = state.V[i]
		}
		state.PC += 2
	case high == 0xF && inst[1] == 0x65:
		for i := 0; i <= int(x); i++ {
			state.V[i] = cpu.memory.Mem[int(state.I)+i]
		}
		state.PC += 2
	}
}
