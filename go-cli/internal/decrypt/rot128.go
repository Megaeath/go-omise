package decrypt

func ROT128(input []byte) []byte {
	output := make([]byte, len(input))
	for i, b := range input {
		output[i] = b ^ 128
	}
	return output
}
