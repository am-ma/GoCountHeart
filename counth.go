package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// しょうせつ
type Novel struct {
	text string
}

// もじすうかぞえまん
type Counter struct {
	novel     Novel
	countChar string
}

// params
type CounthArgs struct {
	filePath  string
	countChar string
}

// constractor ConthArgs
func NewCounthArgs(args []string) *CounthArgs {
	ca := new(CounthArgs)
	ca.filePath = "./novel.text"
	ca.countChar = "♡"

	for i := 0; i < len(args); i++ {
		switch i {
		case 1:
			ca.filePath = os.Args[i]
		case 2:
			ca.countChar = os.Args[i]
		}
	}

	return ca
}

// 文字列の中の文字をかぞえるやつ
func (c Counter) CountCharInText(text string) int {
	chars := []rune(c.novel.text)
	len := len(chars)
	count := 0

	for i := 0; i < len; i++ {
		if string(chars[i]) == c.countChar {
			count++
		}
	}

	return count
}

// もじれつのなかの指定文字を全部数えるやつ
func (c Counter) CountAll() int {
	// TODO: row char count & avg
	return c.CountCharInText(c.novel.text)
}

func main() {
	args := NewCounthArgs(os.Args)

	dat, err := ioutil.ReadFile(args.filePath)
	if err != nil {
		fmt.Printf("ファイルがよめないよ！")
		os.Exit(1)
	}
	text := string(dat)

	novel := Novel{text}
	counter := Counter{novel, args.countChar}

	fmt.Printf("ぜんぶで " + strconv.Itoa(counter.CountAll()) + " 回だよ♡")
}
