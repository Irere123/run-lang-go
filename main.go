package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the RunLang programming language!\n", user.Username)
	fmt.Printf("Free free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
