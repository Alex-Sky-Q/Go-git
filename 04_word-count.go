// Count how many times each word appears in a text

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
		It's my life
		It's now or never
		But I ain't gonna live forever
		I just want to live while I'm alive
		
		My heart is like an open highway
		Like Frankie said, "I did it my way"
		I just want to live while I'm alive
		It's my life`

	textArr := strings.Fields(text)

	words := map[string]int{}

	for _, word := range textArr {
		word = strings.ToLower(word)
		word = strings.TrimSuffix(word, "\"")
		word = strings.TrimPrefix(word, "\"")
		_, ok := words[word]
		if !ok {
			words[word] = 1
		} else {
			words[word]++
		}
	}
	//for w, cnt := range words {
	//	fmt.Printf("The word %q appears %v times\n", w, cnt)
	//}
	fmt.Println(words)
}
