package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"strings"
)

func ScoreStr1(InputStr string) float32 {
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

func BestScore1(buf []float32) byte {
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

func SingleXor1(symbol byte, InputStr []byte) string{

	var result string
	for i := 0; i < len(InputStr); i++{
		result += string(symbol ^ InputStr[i])
	}

	return result
}

func HammingDistance(str1, str2 []byte) int{
	if len(str1) != len(str2){
		panic("Different lengths")
	}

	var res int
	for i := range str1{
		res += bits.OnesCount8(str1[i] ^ str2[i])
	}

	return res
}

func GetLenKey(text string) int {
	temp := text
	KeyLength := 0
	Max := 0
	for i := 1; i < len(text)/2; i++{
		temp = string(temp[len(text) - 1]) + temp[:len(text) - 1]
		Dist := 0
		for j := 0; j < len(text); j++{
			if temp[j] == text[j]{
				Dist += 1
			}
		}
		if Dist > Max{
			Max = Dist;
			KeyLength = i
		}
	}
	return KeyLength
}

func FindSingleXor(InputStr []byte) byte{
	var ans string
	var buf []float32
	for i := 0; i < 256; i++{
		ans = SingleXor1(byte(i), InputStr)
		buf = append(buf, ScoreStr1(ans))
	}

	return BestScore1(buf)
}

func GetKeyText(KeyLength int, text string) string{

	LenBlock := (len(text))/KeyLength	// len blocks
	Block := make([][]byte, KeyLength)

	for i := 0; i < KeyLength; i++{
		Block[i] = make([]byte,LenBlock)
	}

	for i := 0; i < KeyLength; i++{
		for j := 0; j < LenBlock; j++{
			Block[i][j] = text[j*KeyLength + i]
		}
	}

	var key string
	for k := 0; k < KeyLength; k++ {
		key += string(FindSingleXor(Block[k]))
	}

	return key
}

func main(){

	/*file, err := ioutil.ReadFile("task6.txt")

	if err != nil {
		log.Fatal(err)
	}

	text := make([]byte, len(file))

	_, err1 := base64.StdEncoding.Decode(text, file)

	if err1 != nil {
		log.Fatal(err1)
	}

	str1 := string(text)
	KeyVal := GetKeyText(GetLenKey(str1),str1)

	var result []byte
	for i := 0; i < len(text); i++ {
		result = append(result, KeyVal[i % len(KeyVal)] ^ text[i])
	}
	fmt.Println(string(result))*/
	file, err := ioutil.ReadFile("task6.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")

	var text_file string
	for i := 0; i < len(text); i++ {
		text_file += strings.TrimSpace(text[i])
	}

	byteFileText, err1 := base64.StdEncoding.DecodeString(text_file)
	if err1 != nil {
		log.Fatal(err1)
	}

	KeyValue := GetKeyText(GetLenKey(string(byteFileText)), string(byteFileText))

	var result []byte
	for i := 0; i < len(byteFileText); i++ {
		result = append(result, KeyValue[i%len(KeyValue)]^byteFileText[i])
	}

	fmt.Println(string(result))
}

