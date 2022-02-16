package main

import (
	"fmt"

	"github.com/lenimbugua/go-tutorial/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
