package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	buffer := make([]byte, 1024*1024)
	sc.Buffer(buffer, len(buffer))
}

func main() {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	var buildings [4][3][10]int

	for i := 0; i < num; i++ {
		input := toIntArray()
		buildings[input[0]-1][input[1]-1][input[2]-1] += input[3]
	}

	for index, building := range buildings {
		for _, floor := range building {
			for _, room := range floor {
				fmt.Printf(" %v", room)
			}
			fmt.Println()
		}

		if index != 3 {
			fmt.Println("####################")
		}
	}
}

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

// // question_4

// func main() {
// 	sc.Scan()
// 	length, _ := strconv.Atoi(sc.Text())

// 	max := length
// 	min := 0
// 	for i := 0; i < 20; i++ {
// 		index := (max + min) / 2
// 		fmt.Printf("? %v\n", index)
// 		sc.Scan()
// 		input := sc.Text()
// 		if input == "0" {
// 			min = index
// 		} else {
// 			max = index
// 		}

// 		if (max - index) == 1 {
// 			fmt.Printf("! %v\n", index)
// 		}
// 	}
// }

// // question_3
// func main() {
// 	sc.Scan()
// 	sc.Scan()
// 	stringArray := strings.Split(sc.Text(), "")
// 	reversedSlice := make([]string, len(stringArray))
// 	copy(reversedSlice, stringArray)
// 	for i, j := 0, len(reversedSlice)-1; i < j; i, j = i+1, j-1 {
// 		reversedSlice[i], reversedSlice[j] = reversedSlice[j], reversedSlice[i]
// 	}

// 	ans := math.Max(float64(findMaxLength(stringArray)), float64(findMaxLength(reversedSlice)))
// 	if ans == 0 {
// 		fmt.Print(-1)
// 	} else {
// 		fmt.Print(int(ans))
// 	}

// }

// func findMaxLength(stringsArray []string) int {
// 	maxDangoLevel := 0
// 	dangoLevel := 0
// 	for _, input := range stringsArray {

// 		if input == "o" {
// 			dangoLevel++
// 		}

// 		if input == "-" {
// 			if dangoLevel > maxDangoLevel {
// 				maxDangoLevel = dangoLevel
// 			}
// 			dangoLevel = 0
// 		}
// 	}

// 	return maxDangoLevel

// }

// // question_2
// func main() {
// 	sc.Split(bufio.ScanWords)

// 	sc.Scan()
// 	playerNum, _ := strconv.Atoi(sc.Text())

// 	sc.Scan()
// 	color, _ := strconv.Atoi(sc.Text())

// 	c := make([]int, playerNum)
// 	r := make([]int, playerNum)
// 	for i := 0; i < playerNum; i++ {
// 		sc.Scan()
// 		cNum, _ := strconv.Atoi(sc.Text())
// 		c[i] = cNum
// 	}

// 	for i := 0; i < playerNum; i++ {
// 		sc.Scan()
// 		rNum, _ := strconv.Atoi(sc.Text())
// 		r[i] = rNum
// 	}

// 	result := 0
// 	if isColor(c, color) {
// 		result = findMaxPlayer(c, r, color) + 1
// 	} else {
// 		result = findMaxPlayer(c, r, c[0]) + 1
// 	}

// 	fmt.Print(result)

// }

// func isColor(array []int, target int) bool {
// 	result := false
// 	for _, input := range array {
// 		if input == target {
// 			result = true
// 			break
// 		}
// 	}
// 	return result
// }

// func findMaxPlayer(c []int, r []int, color int) int {
// 	result := 0
// 	max := 0

// 	for index, input := range c {
// 		if input == color && max < r[index] {
// 			max = r[index]
// 			result = index
// 		}
// 	}

// 	return result
// }

// question_1
// func main() {
// 	sc.Scan()
// 	length, _ := strconv.Atoi(sc.Text())

// 	sc.Scan()
// 	text := strings.Split(sc.Text(), "")
// 	checkingStatus := "start"
// 	ans := "out"

// 	for i := 0; i < length; i++ {
// 		if text[i] == "|" {
// 			switch checkingStatus {
// 			case "start":
// 				checkingStatus = "checking"
// 			case "checking":
// 				checkingStatus = "finish"
// 			case "finish":
// 				checkingStatus = "finish"
// 			}
// 		}

// 		if text[i] == "*" && checkingStatus == "checking" {
// 			ans = "in"
// 			break
// 		}
// 	}

// 	fmt.Println(ans)

// }

// // スペース区切りで入力をfloat64のスライスに変換する
// func toFloat64Array() []float64 {
// 	var slice []float64
// 	sc.Scan()
// 	x := strings.Split(sc.Text(), " ")
// 	for _, input := range x {
// 		num, _ := strconv.ParseFloat(input, 64)
// 		slice = append(slice, num)
// 	}
// 	return slice
// }

// func toStringArray() []string {
// 	return strings.Split(sc.Text(), " ")
// }
