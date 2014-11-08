// Copyright 2013-2014 Bowery, Inc.

package prompt

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var (
	ErrCTRLC = errors.New("Interrupted (CTRL+C)")
)

const (
	evChar = iota
	evSkip
	evReturn
	evCtrlC
	evBack
	evClear
	evHome
	evEnd
	evRight
	evLeft
	evDel
)

// simplePrompt is a fallback prompt without line editing support.
func (term *Terminal) simplePrompt(prefix string) (string, error) {
	if term.simpleReader == nil {
		term.simpleReader = bufio.NewReader(term.In)
	}

	term.Out.Write([]byte(prefix))
	line, err := term.simpleReader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n ")
	line = strings.TrimLeft(line, " ")

	return line, err
}

// setup initializes a prompt.
func (term *Terminal) setup(buf *Buffer, in io.Reader) (*bufio.Reader, error) {
	cols, _, err := TerminalSize(buf.Out)
	if err != nil {
		return nil, err
	}
	buf.Cols = cols
	input := bufio.NewReader(in)

	err = buf.Refresh()
	if err != nil {
		return nil, err
	}

	return input, nil
}

// read reads a rune and parses ANSI escape sequences found
func (term *Terminal) read(in *bufio.Reader) (int, rune, error) {
	char, _, err := in.ReadRune()
	if err != nil {
		return 0, 0, err
	}

	switch char {
	default:
		// Standard chars.
		return evChar, char, nil
	case tabKey, ctrlA, ctrlB, ctrlE, ctrlF, ctrlG, ctrlH, ctrlJ, ctrlK, ctrlN,
		ctrlO, ctrlP, ctrlQ, ctrlR, ctrlS, ctrlT, ctrlU, ctrlV, ctrlW, ctrlX,
		ctrlY, ctrlZ:
		// Skip.
		return evSkip, char, nil
	case returnKey, ctrlD:
		// End of line.
		return evReturn, char, nil
	case ctrlC:
		// End of line, interrupted.
		return evCtrlC, char, nil
	case backKey:
		// Backspace.
		return evBack, char, nil
	case ctrlL:
		// Clear screen.
		return evClear, char, nil
	case escKey:
		// Functions like arrows, home, etc.
		esc := make([]byte, 2)
		_, err = in.Read(esc)
		if err != nil {
			return -1, char, err
		}

		// Home, end.
		if esc[0] == 'O' {
			switch esc[1] {
			case 'H':
				// Home.
				return evHome, char, nil
			case 'F':
				// End.
				return evEnd, char, nil
			}

			return evSkip, char, nil
		}

		// Arrows, delete, pgup, pgdown, insert.
		if esc[0] == '[' {
			switch esc[1] {
			case 'A', 'B':
				// Up, down.
				return evSkip, char, nil
			case 'C':
				// Right.
				return evRight, char, nil
			case 'D':
				// Left.
				return evLeft, char, nil
			}

			// Delete, pgup, pgdown, insert.
			if esc[1] > '0' && esc[1] < '7' {
				extEsc := make([]byte, 3)
				_, err = in.Read(extEsc)
				if err != nil {
					return -1, char, err
				}

				if extEsc[0] == '~' {
					switch esc[1] {
					case '2', '5', '6':
						// Insert, pgup, pgdown.
						return evSkip, char, err
					case '3':
						// Delete.
						return evDel, char, err
					}
				}
			}
		}
	}

	return evSkip, char, nil
}

// prompt reads from in and parses ANSI escapes writing to buf.
func (term *Terminal) prompt(buf *Buffer, in io.Reader) (string, error) {
	input, err := term.setup(buf, in)
	if err != nil {
		return "", err
	}

	for {
		typ, char, err := term.read(input)
		if err != nil {
			return buf.String(), err
		}

		switch typ {
		case evChar:
			err = buf.Insert(char)
			if err != nil {
				return buf.String(), err
			}
		case evSkip:
			continue
		case evReturn:
			err = buf.EndLine()
			return buf.String(), err
		case evCtrlC:
			err = buf.EndLine()
			if err == nil {
				err = ErrCTRLC
			}

			return buf.String(), err
		case evBack:
			err = buf.DelLeft()
			if err != nil {
				return buf.String(), err
			}
		case evClear:
			err = buf.ClsScreen()
			if err != nil {
				return buf.String(), err
			}
		case evHome:
			err = buf.Start()
			if err != nil {
				return buf.String(), err
			}
		case evEnd:
			err = buf.End()
			if err != nil {
				return buf.String(), err
			}
		case evRight:
			err = buf.Right()
			if err != nil {
				return buf.String(), err
			}
		case evLeft:
			err = buf.Left()
			if err != nil {
				return buf.String(), err
			}
		case evDel:
			err = buf.Del()
			if err != nil {
				return buf.String(), err
			}
		}
	}
}

// password reads from in and parses restricted ANSI escapes writing to buf.
func (term *Terminal) password(buf *Buffer, in io.Reader) (string, error) {
	input, err := term.setup(buf, in)
	if err != nil {
		return "", err
	}

	for {
		typ, char, err := term.read(input)
		if err != nil {
			return buf.String(), err
		}

		switch typ {
		case evChar:
			err = buf.Insert(char)
			if err != nil {
				return buf.String(), err
			}
		case evSkip, evHome, evEnd, evRight, evLeft, evDel:
			continue
		case evReturn:
			err = buf.EndLine()
			return buf.String(), err
		case evCtrlC:
			err = buf.EndLine()
			if err == nil {
				err = ErrCTRLC
			}

			return buf.String(), err
		case evBack:
			err = buf.DelLeft()
			if err != nil {
				return buf.String(), err
			}
		case evClear:
			err = buf.ClsScreen()
			if err != nil {
				return buf.String(), err
			}
		}
	}
}
