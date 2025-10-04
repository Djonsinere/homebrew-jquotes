package src

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"unsafe"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type window struct {
	Row uint16
	Col uint16
}

func check_err(err error) {
	log.Print("[paint]: ", err)
}

func centerText(text string, width int) string {
	textWidth := 0

	for _, char := range text {
		if char > 127 { // Если символ - двойной ширины
			textWidth += 2
		} else {
			textWidth++
		}
	}

	// Определяем, сколько пробелов нужно добавить
	padding := (width - textWidth) / 2
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", padding), text, strings.Repeat(" ", width-textWidth-padding))
}

func Print_square(quote Quote) error {
	terminal_size, err := terminalWindowSize()
	if err != nil {
		check_err(err)
	}

	width := int(terminal_size.Col) - 2

	fmt.Printf("+%s+\n", strings.Repeat("-", width))

	fmt.Printf("|%s|\n", Gray+centerText(quote.Quote_japanese, width)+Reset)

	fmt.Printf("|%s|\n", strings.Repeat(" ", width))

	fmt.Printf("|%s|\n", Green+centerText(quote.Quote_english, width)+Reset)

	fmt.Printf("+%s+\n", strings.Repeat("-", width))

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
