package splitstringintouniquesubstrings

type stringSet map[string]struct{}

func (s stringSet) add(key string)    { s[key] = struct{}{} }
func (s stringSet) remove(key string) { delete(s, key) }
func (s stringSet) contains(key string) bool {
	_, ok := s[key]
	return ok
}

// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description
func maxUniqueSplit(s string) int {
	var res int
	countSubsets(s, 0, 0, stringSet{}, &res)
	return res
}

func countSubsets(s string, iStart, iCurr int, subsets stringSet, res *int) {
	if iCurr == len(s) {
		*res = max(*res, len(subsets))
		return
	}

	// There are 2 seperators before the first letter, and after the last. Hence
	// We cannot "skip adding the sep" if we're at the last char
	if iCurr != len(s)-1 {
		countSubsets(s, iStart, iCurr+1, subsets, res)
	}

	elem := s[iStart : iCurr+1]
	if !subsets.contains(elem) {
		subsets.add(elem)
		countSubsets(s, iCurr+1, iCurr+1, subsets, res)
		subsets.remove(elem)
	}
}
