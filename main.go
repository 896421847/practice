package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(n int, a []int) int {

	if n <= 0 || len(a) == 0 {
		return -1
	}

	sort.Ints(a)

	s := strconv.Itoa(n)
	length := len(s)

	var dfs func(pos int, smaller bool) string

	dfs = func(pos int, smaller bool) string {

		if pos == length {
			return ""
		}

		if smaller {
			max := a[len(a)-1]
			return strconv.Itoa(max) + dfs(pos+1, true)
		}

		limit := int(s[pos] - '0')

		for i := len(a) - 1; i >= 0; i-- {
			d := a[i]

			if pos == 0 && d == 0 {
				continue
			}

			if d > limit {
				continue
			}

			if d < limit {
				return strconv.Itoa(d) + dfs(pos+1, true)
			}

			rest := dfs(pos+1, false)

			if len(rest) == length-pos-1 {
				return strconv.Itoa(d) + rest
			}
		}

		return ""
	}

	result := dfs(0, false)

	if result == "" {
		return largestShorter(a, length-1)
	}

	ans, _ := strconv.Atoi(result)
	return ans
}

func largestShorter(a []int, length int) int {
	if length <= 0 || len(a) == 0 {
		return -1
	}

	max := a[len(a)-1]

	if max == 0 {
		return -1
	}

	result := 0
	for i := 0; i < length; i++ {
		result = result*10 + max
	}

	return result
}

func main() {
	n := 23232
	a := []int{0, 0}

	result := largestNumber(n, a)

	fmt.Println("最大整数为：", result)
}
