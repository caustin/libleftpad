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
