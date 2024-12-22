package core

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
)

const SCREEN_WIDTH = 64
const SCREEN_HEIGHT = 32

type Emulator struct {
	cpu      Cpu
	mem      *Memory
	display  IDisplay
	sound    ISound
	input    IInput
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	running  bool
}

func (emu *Emulator) FillScreenForDebug() {
	emu.display.fillScreen()
}

func NewEmulator() *Emulator {
	display := NewDisplay()
	sound := NewSound()
	memory := NewMemory()
	// Read file and load to memory
	rom, err := os.ReadFile("./roms/testrom/BC_test.ch8")
	//rom, err := os.ReadFile("./roms/demos/Zero Demo [zeroZshadow, 2007].ch8")
	//rom, err := os.ReadFile("./roms/demos/Particle Demo [zeroZshadow, 2008].ch8")
	//rom, err := os.ReadFile("./roms/demos/Stars [Sergey Naydenov, 2010].ch8")
	if err != nil {
		panic("Unable to read rom")
	}
	memory.LoadDisk(rom)
	return &Emulator{
		cpu:     NewCpu(display, sound, memory),
		mem:     memory,
		display: display,
		sound:   sound,
		input:   NewInput(),
	}
}

func (emu *Emulator) Run() {
	fmt.Println("Starting Emulator")
	emu.initialize()
	go emu.cpu.Run()
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
			println("Emulator Stopped")
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
			pixels[offset+0] = 255 * emu.display.getScreen()[x][y] // Blue
			pixels[offset+1] = 255 * emu.display.getScreen()[x][y] // Green
			pixels[offset+2] = 255 * emu.display.getScreen()[x][y] // Red
			pixels[offset+3] = 255                                 // Alpha (fully opaque)
		}
	}

	// Unlock the texture
	emu.texture.Unlock()

	// Clear the renderer and copy the texture to it
	emu.renderer.Clear()
	emu.renderer.Copy(emu.texture, nil, nil)
	emu.renderer.Present()
	// 60 FPS
	sdl.Delay(16)
}
