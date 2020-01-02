package hard

import (
	"testing"
)

func TestGetRainfall(t *testing.T) {
	inputWallArray := []int{0, 0, 0, 0, 0}
	result := GetRainfall(inputWallArray)
	expected := 0
	if result != expected {
		t.Errorf("Failed to satisfy first test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{0, 1, 0, 1}
	result = GetRainfall(inputWallArray)
	expected = 1
	if result != expected {
		t.Errorf("Failed to satisfy second test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{1, 0, 1, 0}
	result = GetRainfall(inputWallArray)
	expected = 1
	if result != expected {
		t.Errorf("Failed to satisfy third test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{1, 2, 3, 4}
	result = GetRainfall(inputWallArray)
	expected = 0
	if result != expected {
		t.Errorf("Failed to satisfy fourth test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{}
	result = GetRainfall(inputWallArray)
	expected = 0
	if result != expected {
		t.Errorf("Failed to satisfy fifth test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{1, 0, 0, 4}
	result = GetRainfall(inputWallArray)
	expected = 2
	if result != expected {
		t.Errorf("Failed to satisfy sixth test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{4, 0, 0, 1}
	result = GetRainfall(inputWallArray)
	expected = 2
	if result != expected {
		t.Errorf("Failed to satisfy sevent test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{4, 0, 0, 4}
	result = GetRainfall(inputWallArray)
	expected = 8
	if result != expected {
		t.Errorf("Failed to satisfy eight test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{1, 0, 3, 4}
	result = GetRainfall(inputWallArray)
	expected = 1
	if result != expected {
		t.Errorf("Failed to satisfy first test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{3, 0, 2, 3}
	result = GetRainfall(inputWallArray)
	expected = 4
	if result != expected {
		t.Errorf("Failed to satisfy first test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}

	inputWallArray = []int{0, 0, 0, 1, 3, 5, 3, 4, 1, 8, 2}
	result = GetRainfall(inputWallArray)
	expected = 7
	if result != expected {
		t.Errorf("Failed to satisfy first test. Input: %v, Result: %v, Expected: %v", inputWallArray, result, expected)
	} else {
		t.Log("Test Passed")
	}
}
