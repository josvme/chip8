package core

type Display struct {
	screen [64][32]byte
}

type IDisplay interface {
	drawSprite(vx int, vy int, N int)
	getScreen() [64][32]byte
}

func NewDisplay() IDisplay {
	return &Display{}
}

func (d *Display) drawSprite(vx int, vy int, N int) {

}

func (d *Display) getScreen() [64][32]byte {
	return d.screen
}
