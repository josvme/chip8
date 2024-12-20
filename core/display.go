package core

type Display struct {
	screen [64][32]byte
}

type IDisplay interface {
	drawSprite(vx int, vy int, N int)
	getScreen() [64][32]byte
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

func (d *Display) drawSprite(vx int, vy int, N int) {

}

func (d *Display) getScreen() [64][32]byte {
	return d.screen
}
