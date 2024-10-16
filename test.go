package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 図形の値を定義
var shapeValues = map[string]int{
	"〇": 0,
	"△": 3,
	"□": 4,
	"☆": 5,
}

// 問題を生成
func generateProblem() (string, string, int) {
	shape1 := "〇"
	shape2 := "△"
	sum := shapeValues[shape1] + shapeValues[shape2]
	return shape1, shape2, sum
}

func main() {
	// 問題を生成
	shape1, shape2, correctSum := generateProblem()

	// 問題を表示
	fmt.Printf("問題: %s %s\n", shape1, shape2)
	fmt.Println("回答を選択してください:")
	fmt.Println("q: !")
	fmt.Println("w: 数字")
	fmt.Println("e: E")

	// ユーザーの回答を取得
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("入力: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// 正解の判定
	if input == "w" && correctSum == shapeValues[shape1]+shapeValues[shape2] {
		fmt.Println("正解！数字です。")
	} else if input == "e" && correctSum != shapeValues[shape1]+shapeValues[shape2] {
		fmt.Println("正解！Eです。")
	} else if input == "q" {
		fmt.Println("正解！Qです。")
	} else {
		fmt.Println("不正解です。")
	}
}
