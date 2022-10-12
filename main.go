package main

import (
	"fmt"
	"sort"
)

func main() {
	var counter = map[string]int{}
	var keys []string
	var input_string string
	_, err := fmt.Scanln(&input_string)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(input_string); i++ {
		if _, ok := counter[string(input_string[i])]; !ok {
			keys = append(keys, string(input_string[i]))
		}
		counter[string(input_string[i])]++
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s%d", key, counter[key])
	}
}
