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
	len_x_quote_eng, len_x_quote_jpn := len(now_quote.Quote_english), len(now_quote.Quote_japanese)

	x := max(len_x_quote_eng, len_x_quote_jpn)
	err = src.Print_square(x, now_quote)
	if err != nil {
		check_err(err)
	}
}
