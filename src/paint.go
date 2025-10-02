package src

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

type window struct {
	Row uint16
	Col uint16
}

func Print_square(len_x int, quote Quote) error {
	fmt.Printf("+%s+\n", strings.Repeat("-", len_x))
	//padding := (len_x - len(quote.Quote_english)) / 2
	fmt.Printf("| %13s", quote.Quote_japanese)
	fmt.Printf("|\n|\n| %13s |\n", quote.Quote_english)

	fmt.Printf("+%s+\n", strings.Repeat("-", len_x))

	return nil
}

func terminalWindowSize() (window, error) {
	win := window{0, 0}
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ), // This parameter, different operating systems may be different, for example: tiocgwinsz_osx
		uintptr(unsafe.Pointer(&win)),
	)
	if int(res) == -1 {
		return window{0, 0}, err
	}

	return win, nil
}
