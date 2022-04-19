package main

import (
	"fmt"
	"time"
)

func main() {
	human1 := [2]string{"sang", "hoon"}
	human2 := [2]string{"hyun", "jae"}

	go humanName(human1)
	humanName(human2)
}

func humanName(names [2]string) {
	for _, name := range names {
		fmt.Println(name)
		time.Sleep(time.Second)
	}
}