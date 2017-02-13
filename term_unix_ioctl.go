// +build linux darwin freebsd openbsd netbsd dragonfly

// Copyright 2013-2015 Bowery, Inc.
// Copyright 2017 Attila Fülöp <attila@fueloep.org>

package prompt

import (
	"os"
	"syscall"
	"unsafe"
	"errors"
)

type Termios syscall.Termios

// setTermios does the system dependent ioctl calls
func getTermios(fd uintptr, req uintptr) (*Termios, error) {
	termios := new(syscall.Termios)
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, req,
		uintptr(unsafe.Pointer(termios)))

	if err != 0 {
		return nil, errors.New(err.Error())
	}
	return (*Termios)(termios), nil
}

// setTermios does the system dependent ioctl calls
func setTermios(fd uintptr, req uintptr, termios *Termios) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, req,
		uintptr(unsafe.Pointer(termios)))

	if err != 0 {
		return errors.New(err.Error())
	}
	return nil
}

// winsize contains the size for the terminal.
type winsize struct {
	rows   uint16
	cols   uint16
	xpixel uint16
	ypixel uint16
}

// TerminalSize retrieves the cols/rows for the terminal connected to out.
func TerminalSize(out *os.File) (int, int, error) {
	ws := new(winsize)

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, out.Fd(),
		uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws)))
	if err != 0 {
		return 0, 0, errors.New(err.Error())
	}
	return int(ws.cols), int(ws.rows), nil
}
