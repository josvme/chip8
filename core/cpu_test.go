package core_test

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	if false {
		t.Errorf("Demo failed")
	}
}

func TestTable(t *testing.T) {
	var tests = []struct{ k, v int }{
		{0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d:%d", tt.k, tt.v)
		t.Run(testname, func(t *testing.T) {
			if tt.k != tt.v {
				t.Errorf("Not equal")
			}
		})
	}
}

func BenchmarkNewCpu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do something
	}
}
