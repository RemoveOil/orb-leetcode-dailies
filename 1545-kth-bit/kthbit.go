package kthbit

const QUESTION_NO = 1545

func FindKthBit(n int, k int) byte {
	remainingLength := (1 << n) - 1
	k--

	var invert byte
	for remainingLength != 1 {
		half := remainingLength / 2
		if half == k {
			return toAsciiByte(1 ^ invert)
		}
		if k > half {
			k = 2*half - k
			invert ^= 1
		}

		remainingLength = half
	}
	return toAsciiByte(invert)
}

func toAsciiByte(val byte) byte {
	return val + byte('0')
}
