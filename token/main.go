package main

import (
	"bufio"
	"fmt"
	"os"

	crypt "github.com/TNCBS/encrypt-decrypt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	encrypted := crypt.EncryptXOR(text, "0E&@w85hetEO7ry6")
	fmt.Println(encrypted)

	decrypted, err := crypt.DecryptXOR(encrypted, "0E&@w85hetEO7ry6")
	if err != nil {
		fmt.Printf("%+v Decrypt Tokens AccountNo failed", err)
	}

	fmt.Println(decrypted)
}
