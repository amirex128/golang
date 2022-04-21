package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	contains := lo.Contains([]string{"amir", "fatemeh"}, "amir")
	fmt.Println(contains)

	sayHello("amir")
}

func sayHello(name string) {
	println("Hello ", name)
}
