package gsm7

const controlChar byte = 0x1B

func mapByteToRune(in map[rune]byte) map[byte]rune {
	out := make(map[byte]rune, len(in))
	for k, v := range in {
		out[v] = k
	}

	return out
}
