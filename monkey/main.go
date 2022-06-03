package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	userObj, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", userObj.Username)
	fmt.Printf("Feel free to type in commands\n")

	tmp := fmt.Sprintf("pwd: %s, Gid: %s, Uid: %s\n", userObj.HomeDir, userObj.Gid, userObj.Uid)
	fmt.Fprintf(os.Stdout, tmp)

	repl.Start(os.Stdin, os.Stdout)
}

//func mytest(in io.Reader, out io.Writer) {
//	scanner := bufio.NewScanner(in)
//
//	for {
//		fmt.Fprintf(out, PROMPT)
//		scanned := scanner.Scan()
//		if !scanned {
//			return
//		}
//
//		line := scanner.Text()
//		l := lexer.New(line)
//
//		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
//			fmt.Fprintf(out, "%+v\n", tok)
//		}
//	}
//}
