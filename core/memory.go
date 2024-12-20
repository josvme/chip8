package core

const loadAddress = 512

type Memory struct {
	Mem [4096]byte
}

func NewMemory() *Memory {
	return &Memory{
		Mem: [4096]byte{},
	}
}

func (m *Memory) LoadDisk(data []byte) {
	for i := 0; i < len(data); i++ {
		m.Mem[loadAddress+i] = data[i]
	}
}
