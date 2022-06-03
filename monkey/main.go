package main

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
	"os"
	"os/user"
)

const PROMPT = ">>"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	tmp := fmt.Sprintf("pwd: %s, Gid: %s, Uid: %s", user.HomeDir, user.Gid, user.Uid)
	fmt.Fprint(os.Stdout, PROMPT)
	fmt.Fprintf(os.Stdout, tmp)

}

func myback() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	mytest(os.Stdin, os.Stdout)
}

func mytest(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
