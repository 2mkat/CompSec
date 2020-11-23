package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func HexToBase(str string) string{
	temp, err := hex.DecodeString(str)
	if err != nil{
		log.Fatal(err)
	}
	return string(base64.StdEncoding.EncodeToString(temp))
}

func main(){
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	fmt.Println(HexToBase(str))
}
