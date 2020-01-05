package hard

import (
	"testing"
)

func TestGetIPAddress(t *testing.T) {
	sampleAddresses := map[string][]string{
		"25525511135": []string{
			"255.255.11.135",
			"255.255.111.35",
		},
		"11111011111": []string{
			"111.110.11.111",
			"111.110.111.11",
		},
	}

	for inputAddress, possibilities := range sampleAddresses {
		result, err := GetIPAddress(inputAddress)
		if err != nil && len(possibilities) != 0 {
			t.Errorf("Expected no error, found error.")
		}
		if len(possibilities) != len(result) {
			t.Errorf("Expected %v results, got %v results. \n Results: %+v", len(possibilities), len(result), result)
		}
		for resultIndex := range result {
			if result[resultIndex] != possibilities[resultIndex] {
				t.Errorf("Incorrect result, expected %v, got %v", possibilities[resultIndex], result[resultIndex])
			}
		}
	}
}

func TestRecursivelyDivideAddresses(t *testing.T) {
	sampleAddress := "123123123123"
	received, err := recursivelyDivideAddresses([]rune(sampleAddress), 0)
	expected := []string{"123.123.123.123"}
	if received == nil {
		t.Errorf("Expected result, got error: %v", err)
	} else {
		if len(*received) != len(expected) {
			t.Errorf("Expected %v result, got %v", len(expected), len(*received))
		}
		for addrIndex := range expected {
			if (*received)[addrIndex] != expected[addrIndex] {
				t.Errorf("Expected %v result, got %v", expected[addrIndex], (*received)[addrIndex])
			}
		}
	}
}
