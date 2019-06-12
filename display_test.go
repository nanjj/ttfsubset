package main

import "testing"

func TestFormatBytes(t *testing.T) {
	if out := FormatBytes([]byte{1, 2}); out != "[]byte{\n1,2,\n}" {
		t.Fatal(out)
	}
	if out := FormatBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
		23, 24, 25}); out != "[]byte{\n1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,\n18,19,20,21,22,23,24,25,\n}" {
		t.Fatal(out)
	}
}
