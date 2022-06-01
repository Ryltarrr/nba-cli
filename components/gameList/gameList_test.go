package gameList

import (
	"testing"
)

var m = New()

func TestShouldScrollDown(t *testing.T) {
	m.viewport.Height = 20
	m.cursor = 10

	// 20/3=6 -> game displayed
	scroll := m.shouldScroll(true)
	if scroll != true {
		t.Fatalf("Should scroll down")
	}
}

func TestShouldScrollUp(t *testing.T) {
	m.viewport.Height = 20
	m.cursor = 1

	// 20/3=6 -> game displayed
	scroll := m.shouldScroll(false)
	if scroll != true {
		t.Fatalf("Should scroll up")
	}
}

func TestShouldNotScrollDown(t *testing.T) {
	m.viewport.Height = 20
	m.cursor = 1

	// 20/3=6 -> game displayed
	scroll := m.shouldScroll(true)
	if scroll != false {
		t.Fatalf("Should not scroll down")
	}
}

func TestShouldNotScrollUp(t *testing.T) {
	m.viewport.Height = 20
	m.cursor = 10

	// 20/3=6 -> game displayed
	scroll := m.shouldScroll(false)
	if scroll != false {
		t.Fatalf("Should not scroll up")
	}
}
