// Given an array of integers that represent the hight of a wall,
// find out how much water will be stored between the wall after a rainfall.

// Example:
// Input: [0,1,0,2,3,0,4,0]
// Output: 4
// Explanation: (1+3)

package hard

import "fmt"

// GetRainfall - Returns the amount of rain stored between the walls
func GetRainfall(wallArray []int) int {
	if len(wallArray) < 3 {
		return 0
	}

	currentMax := 0
	currentMaxIndex := 0
	potentialFill := 0
	currentFill := 0
	for index := range wallArray {
		if index == 0 {
			currentMax = wallArray[index]
			continue
		}
		if currentMax != 0 {
			if wallArray[index-1] < wallArray[index] {
				if wallArray[index] >= currentMax {
					fmt.Println("adding: ", potentialFill)
					currentFill += potentialFill
					potentialFill = 0
				} else {
					fmt.Println("adding: ", potentialFill - (currentMax-wallArray[index])*(index-1-currentMaxIndex))
					currentFill += potentialFill - (currentMax-wallArray[index])*(index-1-currentMaxIndex)
					potentialFill -= (currentMax - wallArray[index]) * (index - 1 - currentMaxIndex)
					potentialFill += currentMax - wallArray[index]
				}
			} else {
				fmt.Println("potentially filling: ", currentMax - wallArray[index])
				potentialFill += currentMax - wallArray[index]
			}
		}
		if currentMax < wallArray[index] {
			fmt.Println("current max update: ", wallArray[index])
			currentMax = wallArray[index]
			currentMaxIndex = index
		}
	}
	return currentFill
}
