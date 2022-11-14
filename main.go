package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		runRepl()
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		fmt.Println("Usage: drafti [script]")
		os.Exit(64)
	}
}

// runFile takes a scripts name as input and starts the interpreter on it.
func runFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("could not open given script '%s'\n", fname)
		os.Exit(65)
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	sc := NewScanner(string(bs))
	ts := sc.Scan()
	fmt.Println(ts)
}

// runRepl starts the REPL for line-by-line interpretation.
func runRepl() {
	// read form standard input
	rd := bufio.NewReader(os.Stdin)
	for {
		// prompt
		fmt.Print("> ")

		// read input
		bs, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// evaluate
		sc := NewScanner(string(append(bs, '\n')))
		ts := sc.Scan()

		// response
		fmt.Println(ts)
	}

}
