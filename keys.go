// Copyright 2013-2014 Bowery, Inc.

package prompt

// Line ending in raw mode.
var crlf = []byte("\r\n")

const (
	tabKey    = '\t'
	backKey   = '\u007f'
	returnKey = '\r'
	escKey    = '\u001B'
	spaceKey  = '\u0020'
)

const (
	ctrlA = iota + 1
	ctrlB
	ctrlC
	ctrlD
	ctrlE
	ctrlF
	ctrlG
	ctrlH
	_ // Same as tabKey.
	ctrlJ
	ctrlK
	ctrlL
	_ // Same as returnKey.
	ctrlN
	ctrlO
	ctrlP
	ctrlQ
	ctrlR
	ctrlS
	ctrlT
	ctrlU
	ctrlV
	ctrlW
	ctrlX
	ctrlY
	ctrlZ
)
