package main

import (
	"fmt"
)


func main() {
	fmt.Println("woot ", arr, str)

	getUniques(str)
}



var arr = [...]int{1, 2, 3, 4, 5, 6}
var str = [...]string{"a", "b", "a", "c", "a", "b"}

func getUniques(s [6]string) {

	uniquemap := make(map[string]string)

	for key, val := range s{
		uniquemap[string(val)] = val
		fmt.Println(key, val)
	}

	fmt.Println(uniquemap)
}
