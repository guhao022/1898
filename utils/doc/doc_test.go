package doc

import "testing"

func Test_Doc(t *testing.T) {
	doc := New("../../dll")

	doc.Analyze().Json()
}


