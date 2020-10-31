package main
import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"log"
)

func main(){
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result, err := hex.DecodeString(str)
	if err != nil{
		log.Fatal(err)
	}
	res := base64.StdEncoding.EncodeToString(result)
	fmt.Println(string(res))
}
