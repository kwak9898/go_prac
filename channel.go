package main

import (
	"fmt"
	"reflect"
)

func isNumbers(numbers int, c chan bool) {
	if reflect.TypeOf(numbers).String() == "int" {
		c <- true
	} else {
		c <- false
	}
}

func main() {
	c := make(chan bool)
	numbers := [4]int{1, 1, 3, 2}
	for _, number := range numbers {
		go isNumbers(number, c)
		fmt.Println(<-c)
	}
}