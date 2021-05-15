package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	var mx int = 0
	var temp string = ""
	var ctr int = 0

	for _, v := range s {

		found := strings.ContainsAny(temp, string(v))

		if !found {
			ctr++
			temp = temp + string(v)
		} else {
			ctr = 1
			temp = "" + string(v)
		}

		if mx < ctr {
			mx = ctr
		}
	}

	return mx
}
