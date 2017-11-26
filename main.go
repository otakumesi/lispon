package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/otakumesi/lispon/lisp"
)

func main() {
	repl()
}

func repl() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for stdin.Scan() {
		fmt.Println("=> ", lisp.Run(stdin.Text()))
		fmt.Print("> ")
	}
}
