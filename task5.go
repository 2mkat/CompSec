package main

import (
	"encoding/hex"
	"fmt"
)

func main()  {

	Str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	KeyValue := "ICE"

	var result []byte
	for i := 0; i < len(Str); i++ {
		result = append(result, KeyValue[i%len(KeyValue)] ^ Str[i])
	}
	fmt.Println(hex.EncodeToString(result))
}




