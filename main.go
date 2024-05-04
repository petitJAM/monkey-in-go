package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! It's Monkey time\n", user.Username)
	fmt.Println("Type in some stuff")

	repl.Start(os.Stdin, os.Stdout)
}
