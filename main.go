package main

import (
	"fmt"
	"os"
	user2 "os/user"
	"seele/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is seele programming language\n", user.Username)
	fmt.Printf("feel free to type in command line\n")
	repl.Start(os.Stdin, os.Stdout)
}
