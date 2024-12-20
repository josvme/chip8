package main

import "chip8/core"

func main() {
	emu := core.NewEmulator()
	emu.FillScreenForDebug()
	emu.Run()
}
