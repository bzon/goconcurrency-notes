package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Bar:", i)
	}
}
