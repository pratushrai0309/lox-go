package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError bool = false

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else if hadError {
		os.Exit(65)
	} else {

		runPrompt()
	}
}

func runFile(fileName string) {

}

func runPrompt() {
	for true {
		fmt.Printf(">")
		var reader = bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		fmt.Println(line)
		run()
		hadError = false
	}
}

func run() {

}

func error(line int, message string) {

}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}
