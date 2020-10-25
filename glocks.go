package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	dat, err := ioutil.ReadFile(file)
	check(err)
	run(string(dat))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		dat, err := reader.ReadBytes('\n')
		check(err)
		run(string(dat))
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
