package main

import (
	"fmt"

	crypt "github.com/TNCBS/encrypt-decrypt"
)

func main() {
	var input string
	var method string
	fmt.Scan(&input)
	fmt.Scan(&method)

	if method == "1" {
		encrypted := crypt.EncryptXOR(input, "0E&@w85hetEO7ry6")
		fmt.Println(encrypted)

		decrypted, err := crypt.DecryptXOR(encrypted, "0E&@w85hetEO7ry6")
		if err != nil {
			fmt.Printf("%+v Decrypt Tokens AccountNo failed", err)
		}

		fmt.Println(decrypted)
	} else if method == "2" {
		decrypted, err := crypt.DecryptXOR(input, "0E&@w85hetEO7ry6")
		if err != nil {
			fmt.Printf("%+v Decrypt Tokens AccountNo failed", err)
		}

		fmt.Println(decrypted)
	}

	// part1 := input[0:3]
	// fmt.Println(part1)
	// part2 := input[len(input)-3 : len(input)]
	// fmt.Println(part2)

	// substring := input[3+1 : len(input)-3+1]
	// fmt.Println(substring)

	// substring = strings.Repeat("X", len(substring))
	// fmt.Println(substring)

	// concat := part1 + substring + part2
	// fmt.Println(concat)

}
