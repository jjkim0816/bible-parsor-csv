package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var path []string

// 성경
type BibleBook struct {
	Title  string
	Phrase []BiblePhrase
}

// 구
type BiblePhrase struct {
	Num    int
	Clause []BibleClause
}

// 절
type BibleClause struct {
	Num     int
	Content string
}

func main() {
	fmt.Println("start bible parsor")
	bible := BibleBook{}

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

		// 첫번째 줄은 제목

	}
}
