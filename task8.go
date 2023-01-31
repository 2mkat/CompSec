package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("task8.txt")
	if err != nil {
		log.Fatal(err)
	}
	filelines := strings.Split(string(file), "\n")

	var (
		nLine     int
		nRepeated int
	)

	for i := 0; i < len(filelines); i++ {
		res, err := hex.DecodeString(strings.TrimSpace(filelines[i]))
		if err != nil {
			log.Fatal(err)
		}

		before := len(res) / 16
		var set = make(map[string]bool)

		for i := 0; i < before; i++ {
			var temp = make([]byte, 16)
			for j := 0; j < 16; j++ {
				temp[j] = res[i*16+j]
			}

			str := string(temp)
			if !set[str] {
				set[str] = true
			}
		}

		if before-len(set) > nRepeated {
			nRepeated = before - len(set)
			nLine = i + 1
		}
	}

	fmt.Println("Detect AES in ECB mode in text line:", nLine, "\nContains duplicate blocks:", nRepeated)
}
