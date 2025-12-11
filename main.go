package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Skarlso/went/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Went programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
