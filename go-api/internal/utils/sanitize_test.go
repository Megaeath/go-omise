package utils

import "testing"

func TestMaskCardNumber(t *testing.T) {
	tests := []struct {
		card     string
		expected string
	}{
		{"4111111111111111", "************1111"},
		{"1234", "1234"},
		{"", ""},
	}

	for _, test := range tests {
		result := MaskCardNumber(test.card)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
