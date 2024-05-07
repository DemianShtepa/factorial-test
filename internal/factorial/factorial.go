package factorial

import (
	"math/big"
)

func Calculate(value uint64) string {
	result := new(big.Int).SetUint64(1)

	for value >= 0 {
		if value == 0 {
			result = result.Mul(result, big.NewInt(1))

			break
		}

		result = result.Mul(result, new(big.Int).SetUint64(value))
		value -= 1
	}

	return result.String()
}
