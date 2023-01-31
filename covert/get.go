package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


func main() {
	var buff []byte
	for i := 0; ; i++ {
		resp, err := http.Get("http://127.0.0.1:1234/?map=" + strconv.Itoa(i))
		if err != nil {
			break
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}
		strBody := string(body)
		startIndex := strings.Index(strBody, "let fieldArr =") + len("let fieldArr =")
		endIndex := startIndex
		for endIndex < len(strBody) && strBody[endIndex] != ']' {
			endIndex++
		}
		arr := strings.Split(strBody[startIndex:endIndex], ",")
		if len(arr) <= 1 {
			break
		}
		for i := 0 ; i < len(arr); i++ {
			var tmp byte
			for j := 7; j >= 0; j-- {
				for i < len(arr) && arr[i] != "1" {
					i++
				}
				if i < len(arr) && arr[i] == "1" && i % 2 == 1{
					tmp |= 1 << j
				}
				i++
			}
			if tmp != 0 {
				buff = append(buff, tmp)
			}
		}
	}
	fmt.Println(string(buff))
}