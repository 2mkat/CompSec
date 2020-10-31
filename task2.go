package main
import(
	"fmt"
	"encoding/hex"
	"log"
)

func main(){
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	res, err1 := hex.DecodeString(str1)
	if err1 != nil {
		log.Fatal(err1)
	}
	res2, err2 := hex.DecodeString(str2)
	if err2 != nil {
		log.Fatal(err2)
	}

	var result string

	for i := 0; i < len(res); i++{
		result += string(res[i] ^ res2[i])
	}

	fmt.Println(hex.EncodeToString([]byte(result)))
}