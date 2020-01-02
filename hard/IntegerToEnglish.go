// Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 231 - 1.

// Example 1:

// Input: 123
// Output: "One Hundred Twenty Three"
// Example 2:

// Input: 12345
// Output: "Twelve Thousand Three Hundred Forty Five"
// Example 3:

// Input: 1234567
// Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
// Example 4:

// Input: 1234567891
// Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

package hard

import ()

const billion = 1000000000
const million = 1000000
const thousand = 1000
const hundred = 100
const ten = 10
const digit = 1

var numMap = map[int]string{
	1: " one ",
	2: " two ",
	3: " three ",
	4: " four ",
	5: " five ",
	6: " six ",
	7: " seven ",
	8: " eight ",
	9: " nine ",
}

func numberToWords(num int) string {
	if num >= billion {
		bdigits := int(num / billion)
		other := num % billion
		return numberToWords(bdigits) + " billion " + numberToWords(other)
	} else if num >= million && num < billion {
		mdigits := int(num / million)
		other := num % million
		return numberToWords(mdigits) + " million " + numberToWords(other)
	} else if num >= thousand && num < million {
		tdigits := int(num / thousand)
		other := num % thousand
		return numberToWords(tdigits) + " thousand " + numberToWords(other)
	} else if num >= hundred && num < thousand {
		hdigits := int(num / hundred)
		other := num % hundred
		return numberToWords(hdigits) + " hundred " + numberToWords(other)
	} else if num >= ten && num < hundred {
		if num > 90 {
			return " ninety " + numberToWords(num-90)
		} else if num > 80 {
			return " eighty " + numberToWords(num-80)
		} else if num > 70 {
			return " seventy " + numberToWords(num-70)
		} else if num > 60 {
			return " sixty " + numberToWords(num-60)
		} else if num > 50 {
			return " fifty " + numberToWords(num-50)
		} else if num > 40 {
			return " fourty " + numberToWords(num-40)
		} else if num > 30 {
			return " thirty " + numberToWords(num-30)
		} else if num > 20 {
			return " twenty " + numberToWords(num-20)
		} else if num > 10 {
			if num == 11 {
				return " eleven "
			} else if num == 12 {
				return " twelve "
			} else if num == 13 {
				return " thirteen "
			} else if num == 14 {
				return " fourteen "
			} else if num == 15 {
				return " fifteen "
			} else if num == 16 {
				return " sixteen "
			} else if num == 17 {
				return " seventeen "
			} else if num == 18 {
				return " eighteen "
			} else {
				return " nineteen "
			}
		}
	}
	return numMap[num]
}
