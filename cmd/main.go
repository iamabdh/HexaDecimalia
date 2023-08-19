package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// command -h <decimal> -> convert into hex number
	// command -d <hex> -> convert into decimal number

	args := os.Args[1:]
	// check number of args have been passed
	if len(args) < 2 {
		fmt.Println("\n\nfor converting decimal numbers to hex: -h <decimal-numbers> \n\nfor converting hex to decimal -d <hex-numbers> ")
		return
	}
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
		decimal := dec(args[1])
		if decimal == nil {
			// handling error
		} else {
			fmt.Println(decimal)
		}
	}
}

func hex(deci int) string {
	hex := strconv.FormatInt(int64(deci), 16)
	return hex
}

func dec(hex string) *big.Int {
	deci := new(big.Int)
	if _, success := deci.SetString(hex, 16); !success {
		return nil
	}
	return deci
}
