package random

import "testing"

// Bytes
func TestString(t *testing.T) {
	str := String(2, 4, []byte("1234567890"))

	println(str)
}

