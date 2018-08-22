package main

import (
	"bufio"
	"fmt"
	"os"
)

// しょうせつ
type Novel struct {
	text string
	rows []string
}

// もじすうかぞえまん
type Counter struct {
	novel     *Novel
	countChar string
}

// params
type CounthArgs struct {
	filePath  string
	countChar string
}

// constractor Novel
func NewNovel(path string) *Novel {
	n := new(Novel)

	text, t := "", ""

	fp, err := os.Open(path)
	if err != nil {
		//    panic(err)
		fmt.Printf("ファイルがよめないよ！")
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		t = scanner.Text()
		n.rows = append(n.rows, t)
		text += t
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	n.text = text

	return n
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
	return c.CountCharInText(c.novel.text)
}

// 行の中の平均をとるやつ
func (c Counter) AvgRows() float64 {
	rowLen := len(c.novel.rows)
	countAll := 0
	// MWMO;
	for i := 0; i < rowLen; i++ {
		countAll += c.CountCharInText(c.novel.rows[i])
	}

	return float64(countAll) / float64(rowLen)
}

func main() {
	args := NewCounthArgs(os.Args)

	novel := NewNovel(args.filePath)
	counter := Counter{novel, args.countChar}

	// TODO: Printf %v
	fmt.Printf("ぜんぶで %d 回だよ♡\n", counter.CountAll())
	fmt.Printf("1行あたりだいたい %.2f 回だよ♡", counter.AvgRows())
}
