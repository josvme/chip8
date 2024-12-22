package main

import "chip8/core"

func main() {
	emu := core.NewEmulator()
	emu.Run()
}
