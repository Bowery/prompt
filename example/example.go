package main

import (
	"fmt"
	"os"

	"github.com/Bowery/prompt"
	"github.com/google/gops/agent"
)

func main() {
	err := agent.Listen(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	terminal, err := prompt.NewTerminal()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer terminal.Close()

	for {
		text, err := terminal.Basic("> ", false)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Println("Typed:", text)
	}
}
