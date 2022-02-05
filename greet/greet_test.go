package greet_test

import (
	"testing"

	"github.com/Sewarion/go-actions-test/greet"
)

func TestGreeting(t *testing.T) {
	wants := "Hello, Julian!"
	got := greet.Greeting("Julian")
	if wants != got {
		t.Errorf("Expected: '%s', got: '%s'", wants, got)
	}

}
