package factorial_test

import (
	"github.com/DemianShtepa/factorial-test/internal/factorial"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		title  string
		input  uint64
		output string
	}{
		{
			title:  "0! = 1",
			input:  0,
			output: "1",
		},
		{
			title:  "8! = 40320",
			input:  8,
			output: "40320",
		},
		{
			title:  "16! = 20922789888000",
			input:  16,
			output: "20922789888000",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.output, factorial.Calculate(c.input))
		})
	}
}
