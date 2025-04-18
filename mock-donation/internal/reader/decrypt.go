package reader

func DecryptROT128(s string) string {
	decoded := make([]rune, len(s))
	for i, r := range s {
		decoded[i] = r - 128
	}
	return string(decoded)
}
