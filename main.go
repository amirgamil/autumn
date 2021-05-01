package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//run stuff
func run(source string) {
	[]Token 
}

//Run the interpreter from a file
func runFile(path string) {
	sourceCode, err := os.Open(path)
	if err != nil {
		fmt.Println("Error trying to open the file! Specifically, ", err)
	}
	byteArray, err := ioutil.ReadAll(sourceCode)
	if err != nil {
		fmt.Println("Error trying to read the file! Specifically, ", err)
	}
	run(string(byteArray))
}

//Run the interpreter in a REPL
func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		run(text)
	}
}

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) > 2 {
		fmt.Println("whatca doing there son")
	} else if len(argsWithProg) == 2 {
		runFile(argsWithProg[1])
	} else {
		//Having a more "intimate" conversation with the interpreter - run an interactive promot (called REPL)
		//called REPL since it goes through a loop of Reading a line of input, Evaluating it, then Printing the result, and Looping back
		runPrompt()
	}
}
