package longesthappystring

import "fmt"

const (
	A           = byte('a')
	B           = byte('b')
	C           = byte('c')
	DUMDUM      = byte('#')
	QUESTION_NO = 1405
)

// https://leetcode.com/problems/longest-happy-string/description/
func LongestHappyString(maxA, maxB, maxC int) string {
	remaining := map[byte]int{A: maxA, B: maxB, C: maxC}
	for k, v := range remaining {
		if v == 0 {
			delete(remaining, k)
		}
	}

	var chars []byte
	for len(remaining) > 0 {
		bannedChar := getNextBannedChar(chars)
		nextChar := getNext(remaining, bannedChar)
		if nextChar == DUMDUM {
			break
		}
		remaining[nextChar]--
		if remaining[nextChar] == 0 {
			delete(remaining, nextChar)
		}
		chars = append(chars, nextChar)
	}

	return string(chars)
}

func getNext(remaining map[byte]int, bannedChar byte) byte {
	res := DUMDUM
	maxRemaining := 0
	for k, v := range remaining {
		if k != bannedChar && v > maxRemaining {
			res = k
			maxRemaining = v
		}
	}
	return res
}

func getNextBannedChar(chars []byte) byte {
	n := len(chars)
	if n < 2 || (chars[n-1] != chars[n-2]) {
		return DUMDUM
	}
	return chars[n-1]
}

func main() {

	fmt.Println(LongestHappyString(1, 1, 7))
	fmt.Println(LongestHappyString(7, 1, 0))
	fmt.Println(LongestHappyString(7, 7, 1))
	fmt.Println(LongestHappyString(0, 0, 1000))
}
