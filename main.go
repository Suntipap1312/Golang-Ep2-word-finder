package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	fmt.Println(words)
	query := os.Args[1:]

queries: // break loop queries
	for _, q := range query { // ค้นหาที่พิมพ์ใน os.args
	search:
		for i, w := range words { // คำที่อยู่ใน corpus
			switch q {
			case "and", "or", "the": // ถ้าเจอ 3 คำนี้ให้ออก loop : search ตัดคำในนี้ออก
				break search
			} // เช็ค query == words(นับทุกตัวรวมถึงตัวซ้ำกันด้วยถ้าไม่มี continue)
			if strings.ToLower(q) == w { // ถ้าไม่มี 3 คำ ข้างบนให้มา condition นี้
				fmt.Printf("#%-2d : %q\n", i+1, w)
				continue queries
				// continue loop queries เริ่มจากด้านบน
			}
		}
	}
}
