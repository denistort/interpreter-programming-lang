package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language!\n", current.Name)
	fmt.Printf("Feel free to type any program / command\n")

	repl.New().Start(os.Stdin)
}
