package hard

import (
	"testing"
)

func Test_IntegerToEnglish(t *testing.T) {
	result := numberToWords(12345678)
	expected := "one"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
