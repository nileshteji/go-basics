package main

import "testing"

// t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("The Length is not valid %v", len(d))

	}
}
