package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		// if output := Add(test.arg1, test.arg2); output != test.expected {
		// 	// t.Errorf("Output %q not equal to expected %q", output, test.expected)
		// 	assert.Equal(t, output, test.expected)
		// }

		assert.Equal(t, Add(test.arg1, test.arg2), test.expected)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
