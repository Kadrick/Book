package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		ok := scanner.Scan()
		if !ok {
			return
		}

		line := scanner.Text()
		tokens := lexer.New(line)

		for tok := tokens.NextToken(); tok.Type != token.EOF; tok = tokens.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
