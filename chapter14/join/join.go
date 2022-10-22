package main

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	}
	if len(phrases) == 1 {
		return phrases[0]
	}
	if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", JoinWithCommas(phrases))

}
