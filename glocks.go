package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lukeomalley/glocks/parsererror"
)

func main() {
	flag.String("file", "", "the script file to execute")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("usage: ./glocks [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(file string) {
	src, err := ioutil.ReadFile(file)
	check(err)
	run(string(src))

	if parsererror.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		src, err := reader.ReadBytes('\n')
		check(err)
		run(string(src))
		parsererror.HadError = false
	}
}

func run(src string) {
	fmt.Println(src)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func handleError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println("[line " + string(line) + "] " + "Error " + where + ": " + message)
	parsererror.HadError = true
}
