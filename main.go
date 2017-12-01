package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/otakumesi/lispon/lisp"
)

const BUFSIZE = 1024

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

	interpreter()
}

func interpreter() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	var sexprs string
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		sexprs = string(buf[:n])
	}
	lisp.Run(sexprs)
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
