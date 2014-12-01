// Copyright 2013-2014 Bowery, Inc.

// Package prompt implements a cross platform line-editing prompt. It also
// provides routines to use ANSI escape sequences across platforms for
// terminal connected io.Readers/io.Writers.
//
// If os.Stdin isn't connected to a terminal or (on Unix)if the terminal
// doesn't support the ANSI escape sequences needed a fallback prompt is
// provided that doesn't do line-editing. Unix terminals that are not supported
// will have the TERM environment variable set to either "dumb" or "cons25".
//
// The keyboard shortcuts are similar to those found in the Readline library:
//
//   - Enter / CTRL+D
//     - End the line.
//   - CTRL+C
//     - End the line, return error `ErrCTRLC`.
//   - Backspace
//     - Remove the character to the left.
//   - CTRL+L
//     - Clear the screen(keeping the current lines content).
//   - Home / End
//     - Jump to the beginning/end of the line.
//   - Left arrow / Right arrow
//     - Move left/right one character.
//   - Delete
//     - Remove the character to the right.
package prompt

import (
	"io"
	"strings"
)

// Basic gets input and if required tests to ensure input was given.
func Basic(prefix string, required bool) (string, error) {
	return Custom(prefix, func(input string) (string, bool) {
		if required && input == "" {
			return "", false
		}

		return input, true
	})
}

// BasicDefault gets input and if empty uses the given default.
func BasicDefault(prefix, def string) (string, error) {
	return Custom(prefix+"(Default: "+def+")", func(input string) (string, bool) {
		if input == "" {
			input = def
		}

		return input, true
	})
}

// Ask gets input and checks if it's truthy or not, and returns that
// in a boolean fashion.
func Ask(question string) (bool, error) {
	input, err := Custom(question+"?(y/n)", func(input string) (string, bool) {
		if input == "" {
			return "", false
		}
		input = strings.ToLower(input)

		if input == "y" || input == "yes" {
			return "yes", true
		}

		return "", true
	})

	ok := false
	if input != "" {
		ok = true
	}

	return ok, err
}

// Custom gets input and calls the given test function with the input to
// check if the input is valid, a true return will return the string.
func Custom(prefix string, test func(string) (string, bool)) (string, error) {
	var err error
	input := ""
	ok := false

	term, err := NewTerminal()
	if err != nil {
		return "", err
	}
	defer term.Close()

	for !ok {
		input, err = term.Prompt(prefix)
		if err != nil && err != io.EOF {
			return "", err
		}

		input, ok = test(input)
	}

	return input, nil
}

// Password retrieves a password from stdin without echoing it.
func Password(prefix string) (string, error) {
	var err error
	input := ""

	term, err := NewTerminal()
	if err != nil {
		return "", err
	}
	defer term.Close()

	for input == "" {
		input, err = term.Password(prefix)
		if err != nil && err != io.EOF {
			return "", err
		}
	}

	return input, nil
}
