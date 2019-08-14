package largest_number_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Sort(Nums(strs))

	result := strings.Join(strs, "")

	var i int
	for i = 0; i < len(result)-1 && result[i] == '0'; i++ {
	}
	return result[i:]
}

type Nums []string

func (ns Nums) Less(i, j int) bool {
	return ns[i]+ns[j] > ns[j]+ns[i]
}
func (ns Nums) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}
func (ns Nums) Len() int {
	return len(ns)
}

func TestLargestNumber(t *testing.T) {
	for i, c := range []struct {
		input  []int
		output string
	}{
		{[]int{0, 0}, "0"},
		{[]int{10, 2}, "210"},
		{nil, ""},
		{[]int{98, 989}, "98998"},
		{[]int{12, 121}, "12121"},
	} {
		assert.Equalf(t, c.output, largestNumber(c.input), "case %d fail", i)
	}
}
