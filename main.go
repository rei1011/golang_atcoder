package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	buffer := make([]byte, 1024*1024)
	sc.Buffer(buffer, len(buffer))
}

func main() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	r := toInt2(sc.Text())
	sc.Scan()
	c := toInt2(sc.Text())

	array := makeIntArray2(c + 1)

	for i := 0; i < r; i++ {
		sum := 0
		for j := 0; j < c; j++ {
			sc.Scan()
			input := toInt2(sc.Text())
			sum += input
			array[j] += input

			fmt.Printf("%v ", input)
			if j == c-1 {
				fmt.Printf("%v\n", sum)
				array[c] += sum
			}
		}
	}

	for index, input := range array {
		if index == len(array)-1 {
			fmt.Printf("%v\n", input)
		} else {
			fmt.Printf("%v ", input)
		}
	}
}

// stringをintに変換する
func toInt2(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

// intの配列を作成
func makeIntArray2(length int) []int {
	return make([]int, length, length)
}
