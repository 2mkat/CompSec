package main
import(
	"fmt"
	"encoding/hex"
	"log"
)

func ScoreStr(InputStr string) float32 {
	characterFreq := map[byte]float32 {
		'a': .08167, 'b': .01492, 'c': .02782, 'd': .04253,
			'e': .12702, 'f': .02228, 'g': .02015, 'h': .06094,
			'i': .06094, 'j': .00153, 'k': .00772, 'l': .04025,
			'm': .02406, 'n': .06749, 'o': .07507, 'p': .01929,
			'q': .00095, 'r': .05987, 's': .06327, 't': .09056,
			'u': .02758, 'v': .00978, 'w': .02360, 'x': .00150,
			'y': .01974, 'z': .00074, ' ': .13000,
	}

	var scores float32
	for i := 0; i < len(InputStr); i++{
		scores += characterFreq[InputStr[i]]
	}
	return scores
}

func BestScore(buf []float32) byte {
	var tp float32
	tp = 0
	var temp byte
	for i := 0; i < 256; i++{
		if buf[i] >= tp{
			tp = buf[i]
			temp = byte(i)
		}
	}
	return temp
}

func SingleXor(symbol byte, InputStr []byte) string{

	var result string
	for i := 0; i < len(InputStr); i++{
		result += string(symbol ^ InputStr[i])
	}

	return result
}


func main(){
	str1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res, err1 := hex.DecodeString(str1)
	if err1 != nil {
		log.Fatal(err1)
	}

	var ans string
	var buf []float32
	for i := 0; i < 256; i++{
		ans = SingleXor(byte(i), res)
		buf = append(buf, ScoreStr(ans))
	}

	fmt.Println(SingleXor(BestScore(buf), res))
}