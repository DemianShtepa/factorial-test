package factorial

func Calculate(value uint64) uint64 {
	var result uint64 = 1

	for value >= 0 {
		if value == 0 {
			result *= 1

			break
		}

		result *= value
		value -= 1
	}

	return result
}
