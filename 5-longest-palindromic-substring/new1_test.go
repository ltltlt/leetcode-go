package lp_test

// O(1) space and O(n^2) time
// simple idea: for each char and each nearby char pair, extend it

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	maxStart, maxEnd := 0, 0

	for i := 0; i < len(s); i++ {
		start, end := extend1(s, i, i)
		if end-start > maxEnd-maxStart {
			maxStart, maxEnd = start, end
		}

		if i < len(s)-1 && s[i] == s[i+1] {
			start, end = extend1(s, i, i+1)
			if end-start > maxEnd-maxStart {
				maxStart, maxEnd = start, end
			}
		}
	}
	return s[maxStart : maxEnd+1]
}

func extend1(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}
