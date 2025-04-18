package repl

import (
	"bufio"
	"fmt"
	"os"

	"grain/evaluator"
	"grain/lexer"
	"grain/object"
	"grain/parser"
)

func Start() {
	env := object.NewEnvironment()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Fprintln(os.Stderr, msg)
			}
			continue
		}

		result := evaluator.Eval(program, env)
		if result != nil {
			fmt.Println(result.Inspect())
		}
	}
}
