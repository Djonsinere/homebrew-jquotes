package main

import (
	"log"

	"github.com/Djonsinere/japanese_quotes.git/src"
)

func check_err(err error) {
	log.Print("[main]: ", err)
}

func main() {

	now_quote, err := src.Quote_rotation()
	if err != nil {
		check_err(err)
	}

	err = src.Print_square(now_quote)
	if err != nil {
		check_err(err)
	}

}
