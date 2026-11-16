package main

import "fmt"

func match(s string, worddic []string) bool {
	words := make(map[string]bool)
	for _, word := range worddic {
		words[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "womendog"
	worddic := []string{"women", "dog"}
	fmt.Println(match(s, worddic))
}
