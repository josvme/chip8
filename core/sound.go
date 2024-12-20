package core

type ISound interface {
	beep()
}
type Sound struct {
}

func NewSound() Sound {
	return Sound{}
}

func (s *Sound) beep() {
	panic("Beep here")
}
