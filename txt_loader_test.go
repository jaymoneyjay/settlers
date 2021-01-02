package main

import "testing"

// TestTxtLoader tests if the number of lines returned by the TxtLoader match
func TestTxtLoader(t *testing.T) {
	want := 43
	loader := NewLoader()
	lines := loader.load("boardAscii.txt")
	numberOfLines := len(lines)

	if numberOfLines != want {
		t.Errorf("Number of lines was incorrect, got: %d, want: %d.", numberOfLines, want)
	}
}
