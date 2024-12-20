package core

type ISound interface {
	beep()
}
type Sound struct {
}

func NewSound() ISound {
	return &Sound{}
}

func (s *Sound) beep() {
	panic("Beep here")
}
