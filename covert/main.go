package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	Map []int
	FieldSize byte
	WillageCount byte
}

func findRandomEmptyEl(arrayMap []int) uint64 {
	size := uint64(len(arrayMap))
	rand.Seed(time.Now().UTC().UnixNano())
	var index uint64
	for {
		index = rand.Uint64() % (size - 1)
		if arrayMap[index] == 0 {
			break
		}
	}
	return index
}

func generateMap(mapIndex uint64, size uint64, willagesCount uint64) []int {

	//engTest := []byte("In computer science, the term artificial intelligence (AI) refers to any human-like intelligence exhibited by a computer, robot, or other machine. In popular usage, artificial intelligence refers to the ability of a computer or machine to mimic the capabilities of the human mind—learning from examples and experience, recognizing objects, understanding and responding to language, making decisions, solving problems—and combining these and other capabilities to perform functions a human might perform, such as greeting a hotel guest or driving a car."+
	//"After decades of being relegated to science fiction, today, AI is part of our everyday lives. The surge in AI development is made possible by the sudden availability of large amounts of data and the corresponding development and wide availability of computer systems that can process all that data faster and more accurately than humans can. AI is completing our words as we type them, providing driving directions when we ask, vacuuming our floors, and recommending what we should buy or binge-watch next. And it’s driving applications—such as medical image analysis—that help skilled professionals do important work faster and with greater success.")
	rusTest := []byte("motematika ybivaet us!")
	msgByte := rusTest
	if mapIndex*willagesCount >= uint64(len(msgByte)*8) {
		return make([]int, 0)
	}
	square := size*size
	Map := make([]int, square)
	block := mapIndex*8
	sp := uint64(0)
	maxStep := square/willagesCount
	rand.Seed(time.Now().UTC().UnixNano())
	for pos := block; pos < block + willagesCount && pos < uint64(len(msgByte)) && willagesCount > 0; pos++ {
		for j := 7; j >= 0 && willagesCount > 0; j-- {
			bit := (msgByte[pos] >> j) & 1
			sp += rand.Uint64() % (3*maxStep/7) + 4*maxStep/7
			if byte(sp) & 1 != bit {
				sp++
			}
			Map[sp] = 1
			willagesCount--
		}
	}
	Map[findRandomEmptyEl(Map)] = 2
	return Map
}

func LoadPage(w http.ResponseWriter, mapIndex uint64) {
	gameMap := generateMap(mapIndex, 32, 64)
	pageTemplate, _ := template.ParseFiles("index.html")
	pageTemplate.Execute(w, Page{gameMap, 32, 64})
}

func router(w http.ResponseWriter, request *http.Request) {
	mapStrIndex, ex := request.URL.Query()["map"]
	if !ex || len(mapStrIndex) == 0 {
		LoadPage(w, 0)
		return
	}
	if ex {
		i, err := strconv.Atoi(mapStrIndex[0])
		if err != nil || i < 0 {
			i = 0
		}
		LoadPage(w, uint64(i))
		return
	}
}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":1234", nil)
}