package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// command -h <decimal> -> convert into hex number
	// command -d <hex> -> convert into decimal number

	args := os.Args[1:]

	if args[0] == "-h" {
		// convert decimal to hex
		strDecimalNumber := args[1]
		intDecimalValue, err := strconv.Atoi(strDecimalNumber)
		if err != nil {
			// handling error
		} else {
			fmt.Println(hex(intDecimalValue))
		}

	} else if args[0] == "-d" {
		// convert hex to decimal
	}
}


func hex (deci int) string {
	hex := strconv.FormatInt(int64(deci), 16)
	return hex
}
