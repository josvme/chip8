package core

import "chip8/utils"

type Display struct {
	screen [64][32]byte
}

type IDisplay interface {
	drawPixels(vx byte, vy byte, pixel byte) byte
	getScreen() [SCREEN_WIDTH][SCREEN_HEIGHT]byte
	clearScreen()
	fillScreen()
}

func NewDisplay() IDisplay {
	return &Display{}
}

func (d *Display) clearScreen() {
	for i := 0; i < len(d.screen); i++ {
		for j := 0; j < len(d.screen[0]); j++ {
			d.screen[i][j] = 0
		}
	}
}

func (d *Display) fillScreen() {
	for i := 0; i < len(d.screen); i++ {
		for j := 0; j < len(d.screen[0]); j++ {
			d.screen[i][j] = 1
		}
	}
}

func (d *Display) drawPixels(vx byte, vy byte, pixel byte) byte {
	vfFlag := uint8(0)
	v := uint8(1)
	for i := 0; i < 8; i++ {
		oldPixel := d.screen[utils.ModFn(int(vx)-i, SCREEN_WIDTH)][utils.ModFn(int(vy), SCREEN_HEIGHT)]
		newPixel := oldPixel ^ (pixel & v)
		d.screen[utils.ModFn(int(vx)-i, SCREEN_WIDTH)][utils.ModFn(int(vy), SCREEN_HEIGHT)] = newPixel
		v = v * 2
		if int8(oldPixel) == 1 && int8(newPixel) == 0 {
			vfFlag = uint8(1)
		}
	}
	return vfFlag
}

func (d *Display) getScreen() [64][32]byte {
	return d.screen
}
