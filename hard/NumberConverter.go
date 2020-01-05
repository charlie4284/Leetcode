package hard

import (
	"errors"
)

// NumberConverter - Convert a Roman numeral into a corresponding decimal value
func NumberConverter(romanNumeral string) (int, error) {
	splitRunes := []rune(romanNumeral)
	prevVal := 0
	totalNum := 0
	sequentialValStack := []int{}
	for runeIndex := range splitRunes {
		if runeIndex == 0 {
			// Base case
			prevVal = toNumber(splitRunes[runeIndex])
			sequentialValStack = append(sequentialValStack, prevVal)
			continue
		}
		currentVal := toNumber(splitRunes[runeIndex])
		if currentVal == prevVal {
			// Same case - append to runestack
			sequentialValStack = append(sequentialValStack, currentVal)
			if len(sequentialValStack) > 3 {
				return 0, errors.New("Invalid string")
			}
		} else if currentVal > prevVal && subtractPairMap[string(splitRunes[runeIndex])] == string(splitRunes[runeIndex-1]) {
			// Prepend case - check stack && subtract prev runestack's value from current rune's value
			valueToSubtract := getTotalValue(sequentialValStack)
			if valueToSubtract == 0 {
				return 0, errors.New("Invalid string")
			}
			valToAdd := currentVal - valueToSubtract
			totalNum += valToAdd
			sequentialValStack = []int{}
		} else if currentVal < prevVal {
			// Append case - add the current value to stack
			prevStack := getTotalValue(sequentialValStack)
			totalNum += prevStack
			sequentialValStack = []int{currentVal}
		} else {
			return 0, errors.New("Invalid string")
		}
		prevVal = currentVal
	}
	totalNum += getTotalValue(sequentialValStack)
	return totalNum, nil
}

func toNumber(romanNum rune) int {
	if val, ok := possibleNumberMap[string(romanNum)]; ok {
		return val
	}
	return 0
}

func getTotalValue(values []int) int {
	totalValue := 0
	for valueIndex := range values {
		totalValue += values[valueIndex]
	}
	return totalValue
}

var subtractPairMap = map[string]string{
	"M": "C",
	"D": "C",
	"C": "X",
	"L": "X",
	"X": "I",
	"V": "I",
}

var possibleNumberMap = map[string]int{
	"I": 1,
	// "IV": 4,
	"V": 5,
	// "IX": 9,
	"X": 10,
	// "XL": 40,
	"L": 50,
	// "XC": 90,
	"C": 100,
	// "CD": 400,
	"D": 500,
	// "CM": 900,
	"M": 1000,
}
