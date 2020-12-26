package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func DecryptAes128Ecb(data []byte, key []byte) string {
	cipher, _ := aes.NewCipher(key)

	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	// remove the padding. The last character in the byte array is the number of padding chars
	paddingSize := int(decrypted[len(decrypted)-1])
	return string(decrypted[0 : len(decrypted)-paddingSize])
}

func main() {
	key := "YELLOW SUBMARINE"

	file, err := ioutil.ReadFile("task7.txt")
	if err != nil {
		log.Fatal(err)
	}
	filelines := strings.Split(string(file), "\n")

	var filetext string
	for i := 0; i < len(filelines); i++ {
		filetext += strings.TrimSpace(filelines[i])
	}

	byteFileText, err := base64.StdEncoding.DecodeString(filetext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(DecryptAes128Ecb(byteFileText, []byte(key)))
}
