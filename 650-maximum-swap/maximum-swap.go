package maximumswap

import (
	"slices"
)

const QUESTION_NO = 650

type e struct{ digit, index int }

type Digits []int

func (d Digits) swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Digits) toInt() int {
	var res int
	for _, digit := range d {
		res = res*10 + digit
	}
	return res
}

func NewDigits(number int) Digits {
	if number == 0 {
		return Digits([]int{0})
	}

	var res []int
	for number > 0 {
		res = append(res, number%10)
		number /= 10
	}
	slices.Reverse(res)
	return Digits(res)
}

// https://leetcode.com/problems/maximum-swap/description/
func MaximumSwap(num int) int {
	digits := NewDigits(num)

	largestOnRight := make([]e, len(digits))
	largestOnRight[len(digits)-1] = e{-1, -1}

	for i := len(digits) - 2; i >= 0; i-- {
		if largestOnRight[i+1].digit >= digits[i+1] {
			largestOnRight[i] = largestOnRight[i+1]
		} else {
			largestOnRight[i] = e{digits[i+1], i + 1}
		}
	}

	for i, digit := range digits {
		lor := largestOnRight[i]
		if digit < lor.digit {
			digits.swap(i, lor.index)
			break
		}
	}
	return digits.toInt()
}
