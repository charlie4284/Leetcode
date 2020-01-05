package hard

import (
	"testing"
)

func TestToNumber(t *testing.T) {
	sampleRead := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for runeVal, number := range sampleRead {
		expected := number
		received := toNumber(runeVal)
		if received != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, received)
		}
	}
}

func TestNumberConvert(t *testing.T) {
	romanToDecMap := map[string]int{
		"I":   1,
		"IV":  4,
		"V":   5,
		"IX":  9,
		"XL":  40,
		"L":   50,
		"XC":  90,
		"C":   100,
		"CD":  400,
		"D":   500,
		"CM":  900,
		"M":   1000,
		"ABC": 0,
	}

	for romanNumeral, decimal := range romanToDecMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}
}

func TestOneToTen(t *testing.T) {
	oneToTenMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	for romanNumeral, decimal := range oneToTenMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}
}

func TestTenToHundred(t *testing.T) {
	tenToHundredMap := map[string]int{
		"XI":    11,
		"XII":   12,
		"XIII":  13,
		"XIV":   14,
		"XV":    15,
		"XVI":   16,
		"XVII":  17,
		"XVIII": 18,
		"XIX":   19,
		"XX":    20,
		"XXX":   30,
		"XL":    40,
		"L":     50,
		"LX":    60,
		"LXX":   70,
		"LXXX":  80,
		"XC":    90,
		"C":     100,
	}
	for romanNumeral, decimal := range tenToHundredMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}
}

func TestHundredToThousand(t *testing.T) {
	hundredToThousandMap := map[string]int{
		"CI":   101,
		"CX":   110,
		"CC":   200,
		"CCC":  300,
		"CD":   400,
		"D":    500,
		"DC":   600,
		"DCC":  700,
		"DCCC": 800,
		"CM":   900,
		"M":    1000,
	}
	for romanNumeral, decimal := range hundredToThousandMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}
}

func TestThousandToMax(t *testing.T) {
	hundredToThousandMap := map[string]int{
		"MI":        1001,
		"MM":        2000,
		"MMM":       3000,
		"MMMI":      3001,
		"MMMCMXCIX": 3999,
	}
	for romanNumeral, decimal := range hundredToThousandMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}

}

func TestRandom(t *testing.T) {
	randomMap := map[string]int{
		"ABC":    0,
		"CXI":    111,
		"CIX":    109,
		"MMCMXI": 2911,
	}
	for romanNumeral, decimal := range randomMap {
		expected := decimal
		result, err := NumberConverter(romanNumeral)
		if result == 0 && err == nil {
			t.Errorf("Incorrect conversion. Expected err, got none.")
		}
		if result != expected {
			t.Errorf("Incorrect conversion. Expected: %v, got: %v", expected, result)
		}
	}
}
