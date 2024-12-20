package core

import "github.com/veandco/go-sdl2/sdl"

type Emulator struct {
	cpu     Cpu
	mem     Memory
	display IDisplay
	sound   Sound
	input   Input
	window  *sdl.Window
	surface *sdl.Surface
	running bool
}

func NewEmulator() *Emulator {
	return &Emulator{
		cpu:     NewCpu(),
		mem:     NewMemory(),
		display: NewDisplay(),
		sound:   NewSound(),
		input:   NewInput(),
	}
}

func (emu *Emulator) Run() {
	emu.initialize()
	for emu.running {
		emu.drawScreen()
		emu.handleInput()
		// play sounds
	}
	emu.window.Destroy()
	sdl.Quit()
}

func (emu *Emulator) initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	emu.window = window
	emu.running = true
}

func (emu *Emulator) handleInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
			println("Quit")
			emu.running = false
			break
		}
	}
}
func (emu *Emulator) drawScreen() {
	surface, err := emu.window.GetSurface()
	if err != nil {
		panic(err)
	}
	emu.surface = surface
	emu.surface.FillRect(nil, 0)
	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(emu.surface.Format, colour.R, colour.G, colour.B, colour.A)
	emu.surface.FillRect(&rect, pixel)
	emu.window.UpdateSurface()

	sdl.Delay(33)
}
