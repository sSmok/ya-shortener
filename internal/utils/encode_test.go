package utils

import (
	"testing"
)

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, ""},
		{1, "b"},
		{62, "ab"},
		{12345, "hnd"},
		{987654321, "zag0eb"},
	}

	for _, test := range tests {
		result := Base62Encode(test.input)
		if result != test.expected {
			t.Errorf("Base62Encode(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
