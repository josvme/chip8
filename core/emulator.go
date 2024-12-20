package core

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const SCREEN_WIDTH = 64
const SCREEN_HEIGHT = 32

type Emulator struct {
	cpu      Cpu
	mem      Memory
	display  IDisplay
	sound    Sound
	input    Input
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	running  bool
}

func (emu *Emulator) FillScreenForDebug() {
	emu.display.fillScreen()
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
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Failed to initialize SDL: %s\n", err)
	}

	// Create a window
	window, err := sdl.CreateWindow("Chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 400, 300, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
	}
	emu.window = window

	// Create a renderer
	renderer, err := sdl.CreateRenderer(emu.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s\n", err)
	}
	emu.renderer = renderer

	// Create a texture to manipulate pixels
	texture, err := emu.renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, SCREEN_WIDTH, SCREEN_HEIGHT)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
	}
	emu.texture = texture
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
	// Lock the texture to directly modify its pixels
	pixels, pitch, err := emu.texture.Lock(nil)
	if err != nil {
		log.Fatalf("Failed to lock texture: %s\n", err)
	}

	// Set random colors for each pixel
	for y := 0; y < SCREEN_HEIGHT; y++ {
		for x := 0; x < SCREEN_WIDTH; x++ {
			offset := y*pitch + x*4 // Calculate the offset in the pixel buffer
			// Generate random ARGB values
			pixels[offset+0] = 0                                   // Blue
			pixels[offset+1] = 255 * emu.display.getScreen()[x][y] // Green
			pixels[offset+2] = 0                                   // Red
			pixels[offset+3] = 255                                 // Alpha (fully opaque)
		}
	}

	// Unlock the texture
	emu.texture.Unlock()

	// Clear the renderer and copy the texture to it
	emu.renderer.Clear()
	emu.renderer.Copy(emu.texture, nil, nil)
	emu.renderer.Present()
	// 30 FPS
	sdl.Delay(33)
}
