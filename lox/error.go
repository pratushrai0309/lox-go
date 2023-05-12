package lox

import "fmt"

func error(line int, message string) {

}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}
