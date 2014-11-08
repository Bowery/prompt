// Copyright 2013-2014 Bowery, Inc.

package prompt

import (
	"syscall"
)

const (
	tcgets  = syscall.TCGETS
	tcsets  = syscall.TCSETS
	tcsetsf = 0x5404
)
