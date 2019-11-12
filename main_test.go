package main

import (
	"testing"
)

type (
	doTestInfo struct {
		in  string
		out int64
	}
)

var (
	doTestValues = []doTestInfo{
		doTestInfo{"0xFF", 255},
		doTestInfo{"FF", 255},
		doTestInfo{"A", 10},
		doTestInfo{"", 0},
	}
)

func TestDo(t *testing.T) {
	for _, d := range doTestValues {
		// Arrange
		// Act
		got, err := Do(d.in)

		// Assert
		if err != nil {
			t.Error(err)
		}

		if d.out != got {
			t.Errorf("want: %v, but got:%v\n", d.out, got)
		}
	}
}
