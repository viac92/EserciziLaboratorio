package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string
	mymap := make(map[byte]int)

	fmt.Scan(&word)

	f(word, mymap)

	for key, value := range mymap {
		fmt.Println(string(key), ":", value)
	}
}

func f(s string, mappa map[byte]int) {
	if len(s) == 0 {
		return
	}
	r := s[0]
	if strings.IndexByte("aeiouAEIOU", r) != -1 {
		mappa[r]++
	}
	
	f(s[1:], mappa)
}