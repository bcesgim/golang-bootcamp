/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	substrs := strings.Fields(s)
	fmt.Println(substrs)

	// key,value  ==> word,count
	// iterate over substrs, if key in map, value++, else create

	counts := make(map[string]int)

	for _, word := range substrs {
		_, ok := counts[word]

		if ok == true {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}

	return counts
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	//wc.Test(WordCount)
}
