package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// スペース区切りで入力をintのスライスに変換する
func toIntArray() []int {
	var slice []int
	sc.Scan()
	x := strings.Split(sc.Text(), " ")
	for _, input := range x {
		num, _ := strconv.Atoi(input)
		slice = append(slice, num)
	}
	return slice
}

// スペース区切りで入力をfloat64のスライスに変換する
func toFloat64Array() []float64 {
	var slice []float64
	sc.Scan()
	x := strings.Split(sc.Text(), " ")
	for _, input := range x {
		num, _ := strconv.ParseFloat(input, 64)
		slice = append(slice, num)
	}
	return slice
}

func toStringArray() []string {
	return strings.Split(sc.Text(), " ")
}

// 空白もしくは改行ごとにinputを読み込むための設定
func scanWords() {
	sc.Split(bufio.ScanWords)
}

// stringをintに変換する
func toInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

// intの配列を作成
func makeIntArray(length int) []int {
	return make([]int, length, length)
}

// intの二次元配列を作成
func makeIntTwoDimensionalArray(n int, m int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
	}
	return graph
}

// 正規表現にマッチするか判定する
func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

// runeからintに変換する
func runeToInt(input rune) int {
	return int(input - '0')
}

// alphabetを順に出力する
func printAlphabet() {
	for i := 0; i < 26; i++ {
		fmt.Println(string('a' + i))
	}
}

// 文字列を逆順にするための関数
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// intの配列を降順でsortする
func sortIntReverse(array []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
}

// intの絶対値を取得する関数
func intAbs(input int) int {
	if input >= 0 {
		return input
	} else {
		return -input
	}
}
