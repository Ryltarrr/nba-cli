package menu

import "testing"

func TestShouldAutoComplete(t *testing.T) {
	newVal, newPos := autoComplete("test")
	if newVal != "test-" {
		t.Fatalf("Should autocomplete")
	}

	if newPos != 6 {
		t.Fatalf("Cursor should be after autocompleted character")
	}
}

func TestShouldNotAutoComplete(t *testing.T) {
	newVal, newPos := autoComplete("t")
	if newVal != "t" {
		t.Fatalf("Should not auto-complete")
	}
	if newPos != 1 {
		t.Fatalf("Wrong cursor placement")
	}
}

func TestShouldAutoDelete(t *testing.T) {
	newVal := autoDelete("test-")
	if newVal != "tes" {
		t.Fatalf("Should autodelete")
	}
}

func TestShouldNotAutoDelete(t *testing.T) {
	newVal := autoDelete("test")
	if newVal != "test" {
		t.Fatalf("Should not autodelete")
	}
}
