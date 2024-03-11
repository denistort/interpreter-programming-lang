package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
	"os"
	"strings"
)

const Prompt = "-> "

type Repl struct {
}

func New() *Repl {
	return &Repl{}
}
func (repl *Repl) Start(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "quit" {
			break
		}
		if strings.HasPrefix(line, "build") {
			repl.runFromFile(line)
			break
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func (repl *Repl) runFromFile(line string) {
	split := strings.Split(line, "build")
	clearPath := strings.TrimSpace(strings.Join(split, ""))

	file, err := os.ReadFile(clearPath)
	if err != nil {
		panic(err)
	}
	l := lexer.New(string(file))
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
