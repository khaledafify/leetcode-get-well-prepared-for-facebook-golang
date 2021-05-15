package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	var mx int = 0
	var temp string = ""

	for _, v := range s {

		found := strings.Index(temp, string(v))

		if found == -1 {
			temp = temp + string(v)
		} else {
			temp = temp[found+1:] + string(v)
		}

		if mx < len(temp) {
			mx = len(temp)
		}
	}

	return mx
}
