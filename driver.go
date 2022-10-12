package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"github.com/pkg/term"
	"log"
	"os"
)

func main() {

	plainText, err := os.ReadFile("input.txt")
	check(err)

	cipherKey := getKey()
	encryptedHexText := encryptAES([]byte(cipherKey), plainText)

	fo, err := os.Create("output.txt")
	check(err)
	defer func() {
		err = fo.Close()
		check(err)
	}()

	_, err = fo.Write([]byte(encryptedHexText))
	check(err)

	_ = decryptAES([]byte(cipherKey), encryptedHexText)

	genericLog(plainText, cipherKey, encryptedHexText)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCh() []byte {
	t, _ := term.Open("/dev/tty")
	_ = term.RawMode(t)
	bites := make([]byte, 3)
	numRead, _ := t.Read(bites)
	_ = t.Restore()
	_ = t.Close()
	return bites[0:numRead]
}

func getKey() string {
	fmt.Print("Enter a 32-bit cipher key.\n[ 0] >: ")

	var password string
	var star string

	for {

		password = ""
		star = ""

		for {

			var x = getCh()

			if bytes.Compare(x, []byte{13}) == 0 {
				break
			}

			password = password + string(x)
			star = star + "*"
			fmt.Printf("\r[%2d] >: %s", len(password), star)
		}

		if len(password) == 32 {
			break
		}

		fmt.Print("\n\nValue was not a 32-bit key. Try again!\n>: ")
	}

	fmt.Println()

	return password
}

func encryptAES(cipherKey []byte, plainText []byte) string {

	cipher, err := aes.NewCipher(cipherKey)
	check(err)

	output := make([]byte, len(plainText))

	cipher.Encrypt(output, plainText)

	return hex.EncodeToString(output)
}

func decryptAES(cipherKey []byte, ct string) string {
	cipherText, _ := hex.DecodeString(ct)

	cipher, err := aes.NewCipher(cipherKey)
	check(err)

	plainText := make([]byte, len(cipherText))
	cipher.Decrypt(plainText, cipherText)

	return string(plainText[:])
}

func genericLog(plainText []byte, cipherKey string, encryptedHexText string) {
	fl, err := os.OpenFile("midput.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	check(err)
	defer func(f *os.File) {
		err := f.Close()
		check(err)
	}(fl)

	log.SetOutput(fl)
	log.Printf("\n\nInput:\n%s\n\nKey:\n%s\n\nOutput:\n%s\n", plainText, cipherKey, encryptedHexText)
}
