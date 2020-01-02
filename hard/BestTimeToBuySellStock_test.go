package hard

import (
	"testing"
)

func TestBestTimeToBuySellStock(t *testing.T) {
	result := maxProfit(2, []int{2,4,1})
	expected := 2
	if result != expected {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}

	result = maxProfit(2, []int{3,2,6,5,0,3})
	expected = 7
	if result != expected {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}