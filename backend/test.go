package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 図形の値を定義
var shapeValues = map[string]int{
	"〇": 0,
	"△": 3,
	"□": 4,
	"☆": 5,
}

// ランダムな問題を生成
func generateProblem() (string, string, int) {
	shapes := []string{"〇", "△", "□", "☆"} // 図形のリスト

	// ランダムに2つの図形を選択
	shape1 := shapes[rand.Intn(len(shapes))]
	shape2 := shapes[rand.Intn(len(shapes))]

	// 合計を計算
	sum := shapeValues[shape1] + shapeValues[shape2]
	return shape1, shape2, sum
}

// 数字の選択肢をランダム生成
func random() (int, int) {
	result := rand.Intn(2)
	option := rand.Intn(11)
	return result, option
}

// オプションを生成
func generateOptions(correctSum int) int {
	rnd, option := random()
	if rnd == 0 {
		return correctSum
	} else {
		return option
	}
}

// 10%の確率で「Q」を付けるかどうかを判定する関数
func shouldAddQ() bool {
	return rand.Float64() < 0.1
}

func main() {
	// ランダムシードを設定
	rand.Seed(time.Now().UnixNano())

	// 10問のループ
	correctAnswernum := 0
	falseAnswernum := 0

	for i := 1; i <= 10; i++ {
		// 問題を生成
		shape1, shape2, correctSum := generateProblem()
		option := generateOptions(correctSum)

		// 10%の確率で「Q」を付ける
		withQ := shouldAddQ()
		questionPrefix := ""
		if withQ {
			questionPrefix = "Q"
		}

		// 問題を表示
		fmt.Printf("\n問題 %d: %s%s %s\n", i, questionPrefix, shape1, shape2)
		fmt.Println("q: !")
		fmt.Println("w:", option)
		fmt.Println("e: E")

		// ユーザーの回答を取得
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("入力: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 正解の判定
		if withQ && input == "q" {
			fmt.Println("正解。!です。")
			correctAnswernum++
		} else if !withQ && input == "w" {
			if option == correctSum {
				fmt.Println("正解。", correctSum, "です。")
				correctAnswernum++
			} else {
				fmt.Println("不正解です。正解は", correctSum, "です。")
				falseAnswernum++
			}
		} else if !withQ && input == "e" && option != correctSum {
			fmt.Println("正解。Eです。")
			correctAnswernum++
		} else {
			fmt.Println("不正解です。")
			falseAnswernum++
		}
	}

	fmt.Printf("\n10問が終了しました。お疲れ様でした！\n正解数: %d, 不正解数: %d\n", correctAnswernum, falseAnswernum)
}
