package leftpad

import "testing"

func TestLength(t *testing.T) {
	expected := "*********I"
	expectedLen := len(expected)

	res := PadLeft("I", "*", 10)
	if len(res) != len(res) {
		t.Errorf("Expected %d, got %d", expectedLen, len(res))
	}
}

func TestPadLeft(t *testing.T) {
	expected := "*********I"

	result := PadLeft("I", "*", 10)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
