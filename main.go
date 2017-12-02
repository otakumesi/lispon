package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/otakumesi/lispon/lisp"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
  %s FILE_NAME
      Run the file with lispon
  %s -repl
      Start up REPL
`, os.Args[0], os.Args[0], os.Args[0])
	}
	replFlag := flag.Bool("repl", false, "Start up REPL")
	flag.Parse()

	if *replFlag {
		repl()
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	lisp.Interpreter(os.Args[1])
}

func repl() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for stdin.Scan() {
		for _, result := range lisp.Run(stdin.Text()) {
			fmt.Println("=> ", result)
		}
		fmt.Print("> ")
	}
}
