package leftpad

import "testing"

func TestLength(t *testing.T) {
	expected := "*********I"
	expectedLen := len(expected)

	res, err := PadLeft("I", "*", 10)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if len(res) != len(res) {
		t.Errorf("Expected %d, got %d", expectedLen, len(res))
	}
}

func TestPadLeft(t *testing.T) {
	expected := "*********I"

	result, err := PadLeft("I", "*", 10)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBadTargetLength(t *testing.T) {
	padLengths := [3]int{3, 2, 1}
	for _, length := range padLengths {
		_, err := PadLeft("src", "*", length)
		if err == nil {
			t.Error("Expected length error, got nil")
		}
	}
}

func TestBadPaddingChar(t *testing.T) {
	padChars := [3]string{"", "aa", "  "}

	for _, pchar := range padChars {
		_, err := PadLeft("src", pchar, 10)
		if err == nil {
			t.Error("Expected padding char error, got nil")
		}
	}
}
