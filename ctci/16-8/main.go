package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var number string
	if len(os.Args) > 1 {
		number = os.Args[1]
	}
	output := ""

	numbers := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
	}

	end := len(number)

	if end >= 4 {
		idx := number[end-4 : end-3]
		num, ok := numbers[idx]
		if !ok {
			log.Fatal("malformed number")
		}
		output += num + " " + "thousand"
	}

	if end >= 3 {
		idx := number[end-3 : end-2]
		num, ok := numbers[idx]
		if !ok {
			log.Fatal("malformed number")
		}
		output += " " + num + " " + "hundret"
	}

	subhund := ""
	tens := map[string]string{
		"1": "eleven",
		"2": "twelve",
		"3": "thirteen",
		"4": "fourteen",
		"5": "fifteen",
		"6": "sixteen",
		"7": "seventeen",
		"8": "eighteen",
		"9": "nineteen",
	}
	tenners := map[string]string{
		"0": "",
		"2": "twenty",
		"3": "thirty",
		"4": "fourty",
		"5": "fifty",
		"6": "sixty",
		"7": "seventy",
		"8": "eighty",
		"9": "ninety",
	}
	if end > 1 {
		if number[end-2:end-1] == "1" {
			subhund = " " + tens[number[end-1:end]]
		} else {
			fmt.Println(number[end-2 : end-1])
			ten, ok := tenners[number[end-2:end-1]]
			if !ok {
				log.Fatal("malformed number")
			}
			subhund += " " + ten
		}
		if number[end-2:end-1] != "1" {
			subhund += " " + numbers[number[end-1:end]]
		}
	}
	fmt.Println(output + subhund)
}
