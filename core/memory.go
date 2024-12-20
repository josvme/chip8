package core

type Memory struct {
	Mem          [4096]byte
	Start        int
	ProgramStart int
}

func NewMemory() Memory {
	return Memory{
		Mem:          [4096]byte{},
		Start:        0,
		ProgramStart: 512,
	}
}
