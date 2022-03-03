package main

import (
	"fmt"
	"monkey01/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programing language!\n", user.Username)
	fmt.Printf("Fell free to type in command\n")

	repl.Start(os.Stdin, os.Stdout)
}
