package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func ScoreStr_(InputStr string) float32 {
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

func BestScore_(buf []float32) byte {
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

func SingleXor_(symbol byte, InputStr []byte) string{

	var result string
	for i := 0; i < len(InputStr); i++{
		result += string(symbol ^ InputStr[i])
	}

	return result
}


func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var resultStr []string
	var buf2 []float32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str1 := scanner.Text()
		res, err1 := hex.DecodeString(str1)

		if err1 != nil {
			log.Fatal(err1)
		}

		var ans string
		var buf []float32
		for i := 0; i < 256; i++ {
			ans = SingleXor_(byte(i), res)
			buf = append(buf, ScoreStr_(ans))
		}
		resultStr = append(resultStr, SingleXor_(BestScore_(buf), res))
		buf2 = append(buf2, ScoreStr_(SingleXor_(BestScore_(buf), res)))
	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultStr[BestScore_(buf2)])
}