package hard

import (
	"errors"
	"fmt"
)

// GetIPAddress - Gets all possible IP Address combinations
func GetIPAddress(address string) ([]string, error) {
	possibleAddresses := []string{}
	if !checkValidLength(address) {
		return possibleAddresses, errors.New("Invalid address length")
	}
	adderesses, err := recursivelyDivideAddresses([]rune(address), 0)
	if err != nil {
		return possibleAddresses, err
	}
	if adderesses == nil || len(*adderesses) == 0 {
		return possibleAddresses, errors.New("No valid addresses found")
	}
	return *adderesses, nil
}

func checkValidLength(address string) bool {
	if len(address) < 4 {
		return false
	}
	return true
}

func recursivelyDivideAddresses(addressRunes []rune, level int) (*[]string, error) {
	// Fail case: address starts with 0
	if len(addressRunes) == 0 {
		return nil, fmt.Errorf("No rune providied in level: %v", level)
	} else if level == 3 && len(addressRunes) > 3 {
		return nil, fmt.Errorf("More than 3 runes providied in level: %v", level)
	} else if level == 3 {
		finalAddress := ""
		for runeIndex := range addressRunes {
			finalAddress += string(addressRunes[runeIndex])
		}
		fmt.Println(finalAddress)
		return &[]string{finalAddress}, nil
	}

	possibleAddresses := []string{}
	singleDigit := 0
	if len(addressRunes) >= 1 {
		//x.~
		singleDigit = int(addressRunes[0] - '0')
		if nextLevelAddresses, err := recursivelyDivideAddresses(addressRunes[1:], level+1); err == nil && nextLevelAddresses != nil && len(*nextLevelAddresses) > 0 {
			currentLevelAddress := string(addressRunes[0]) + "."
			fmt.Println(currentLevelAddress)
			possibleAddresses = append(possibleAddresses, appendAddresses(currentLevelAddress, nextLevelAddresses)...)
		}
	}

	if singleDigit != 0 && len(addressRunes) >= 2 {
		//xx.~
		if nextLevelAddresses, err := recursivelyDivideAddresses(addressRunes[2:], level+1); err == nil && nextLevelAddresses != nil && len(*nextLevelAddresses) != 0 {
			currentLevelAddress := string(addressRunes[0]) + string(addressRunes[1]) + "."
			fmt.Println(currentLevelAddress)
			possibleAddresses = append(possibleAddresses, appendAddresses(currentLevelAddress, nextLevelAddresses)...)
		}
	}

	if singleDigit != 0 && len(addressRunes) >= 3 {
		tripleDigit := int(addressRunes[0]-'0')*100 + int(addressRunes[1]-'0')*10 + int(addressRunes[2]-'0')
		if tripleDigit < 256 {
			if nextLevelAddresses, err := recursivelyDivideAddresses(addressRunes[3:], level+1); err == nil && nextLevelAddresses != nil && len(*nextLevelAddresses) != 0 {
				currentLevelAddress := string(addressRunes[0]) + string(addressRunes[1]) + string(addressRunes[2]) + "."
				fmt.Println(currentLevelAddress)
				possibleAddresses = append(possibleAddresses, appendAddresses(currentLevelAddress, nextLevelAddresses)...)
			}
		}
	}
	fmt.Println("level: ", level, "single digit: ", singleDigit, "len: ", len(addressRunes))
	if len(possibleAddresses) > 0 {
		return &possibleAddresses, nil
	}
	return nil, fmt.Errorf("No possible addresses found. Level: %v", level)
}

func appendAddresses(baseAddr string, addrList *[]string) []string {
	appendedAddresses := []string{}
	if addrList == nil || len(*addrList) == 0 {
		return appendedAddresses
	}
	for addrIndex := range *addrList {
		newAddress := baseAddr + (*addrList)[addrIndex]
		appendedAddresses = append(appendedAddresses, newAddress)
	}
	return appendedAddresses
}
