package core

type IInput interface {
	isKeyPressed(key byte) bool
	scanKey() byte
}
type Input struct {
	keysPressed map[byte]bool
}

func (inp *Input) isKeyPressed(key byte) bool {
	return inp.keysPressed[key]
}

func (inp *Input) scanKey() byte {
	for {
		for k, v := range inp.keysPressed {
			if v {
				return k
			}
		}
	}
}

func NewInput() IInput {
	m := make(map[byte]bool)
	for i := 0; i <= 16; i++ {
		m[byte(i)] = false
	}
	return &Input{
		keysPressed: m,
	}
}
