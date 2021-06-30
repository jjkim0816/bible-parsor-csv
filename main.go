package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var path []string

// 성경
type BibleBook struct {
	Title  string
	Phrase []BibleChapter
}

// 구
type BibleChapter struct {
	Num    int
	Clause []BibleVerse
}

// 절
type BibleVerse struct {
	Num     int
	Content string
}

// method
// bible
func (bb *BibleBook) SetTitle(data []byte) {
	bb.Title = string(data)
}

// chapter
func (bc *BibleChapter) SetChpaterNum(data []byte) {
	bc.Num, _ = strconv.Atoi(string(data))
}

func (bc *BibleChapter) SetChapter(data BibleVerse) {
	bc.Clause = append(bc.Clause, data)
}

// verse
func (bv *BibleVerse) SetVerseNum(data []byte) {
	bv.Num, _ = strconv.Atoi(string(data))
}

func (bv *BibleVerse) SetVerse(data []byte) {
	bv.Content = string(data)
}

func main() {
	fmt.Println("start bible parsor")

	book := BibleBook{}
	// chapter := BibleChapter{}
	// verse := BibleVerse{}
	first := false

	root, _ := os.Getwd()
	path = append(path, root)
	path = append(path, "/files/구약/우리말4판_OT01창세기.txt")
	fmt.Printf("path : %s\n\n", path)
	fileUrl := strings.Join(path, "")

	open, err := os.Open(fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer open.Close()

	reader := bufio.NewReader(open)
	for {
		time.Sleep(time.Second * time.Duration(1))

		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			log.Fatal("isPrefix : ", isPrefix, ", err : ", err)
		}

		// text : []bytes
		// n : bytes length
		// err : error
		text, n, err := transform.Bytes(korean.EUCKR.NewDecoder(), line)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("text [%d]: %s\n", n, string(text))
		fmt.Printf("text [%d]: %v\n", n, text)

		// 값이 없는 경우 continue
		if len(text) == 0 {
			fmt.Println("len(text): '", len(text), "' is skip...")
			continue
		}

		// 첫번째 줄은 제목
		if !first {
			book.SetTitle(text)
			first = true
			continue
		}

		// 절이 아닌 경우 continue
		if text[0] < 48 || text[0] > 57 {
			fmt.Println("text[0] : '", text[0], "' is skip...")
			continue
		}

		//
	}
}
