package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yellowpuki/codepractice/base"
)

func TestFibo(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		out  int
	}{
		{
			name: "Case with zero value",
			in:   0,
			out:  0,
		},
		{
			name: "Case with 1",
			in:   1,
			out:  1,
		},
		{
			name: "Case with 5",
			in:   5,
			out:  8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.out, base.Fibo(tc.in))
		})
	}
}
