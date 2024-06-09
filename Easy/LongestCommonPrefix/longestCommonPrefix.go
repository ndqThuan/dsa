package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"aaa", "aa", "aaa"}
	slices.Sort(strs)

	prefix := ""
	firstString := strs[0]
	lastString := strs[len(strs)-1]

	for i := 0; i < min(len(firstString), len(lastString)); i++ {
		if firstString[i] == lastString[i] {
			prefix += string(firstString[i])
		} else {
			break
		}
	}

	fmt.Println(prefix)
}
