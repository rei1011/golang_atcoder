package main

import (
	"bufio"
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
