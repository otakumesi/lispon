package main

import (
	"bufio"
	"fmt"
	"os"

	"./lisp"
)

func main() {
	repl()
}

func repl() {
	lisp.Init()
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for stdin.Scan() {
		fmt.Println("=> ", lisp.Run(stdin.Text()))
		fmt.Print("> ")
	}
}
