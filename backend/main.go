package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 図形の値を定義
var shapeValues = map[string]int{
	"〇": 0,
	"△": 3,
	"□": 4,
	"☆": 5,
}

// 問題を格納する構造体
type Problem struct {
	Shape1     string
	Shape2     string
	correctSum int
	option 	   int
	WithQ      bool
}

var currentProblem Problem

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

// 選択肢を生成
func generateOptions(correctSum int) int {
	rnd := rand.Intn(2)
	option := rand.Intn(11)
	if rnd == 0 {
		return correctSum
	} else {
		return option
	}
}

// 10%の確率で「Q」を付けるかどうかを判定する関数
func shouldAddQ() (string, bool) {
	questionPrefix := ""
	withQ := false
	if rand.Float64() < 0.1 {
		questionPrefix = "Q"
		withQ = true
	}
	return questionPrefix, withQ
}

// 正解の判定
func judgement(input string, withQ bool, option int, correctSum int) (int, int) {
	correctAnswernum := 0
	falseAnswernum := 0

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
	} else if !withQ && input == "e" {
		if option != correctSum {
			fmt.Println("正解。Eです。")
			correctAnswernum++
		} else {
			fmt.Println("不正解です。正解は", correctSum, "です。")
			falseAnswernum++
		}
	} else {
		fmt.Println("不正解です。")
		falseAnswernum++
	}
	return correctAnswernum, falseAnswernum
}

func main() {
	// エンジン作成
	engine := gin.Default()

	// CORSの許可
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 問題提供
	engine.GET("/problem", func(c *gin.Context) {
		shape1, shape2, correctSum := generateProblem()
		option :=generateOptions(correctSum)
		_, withQ := shouldAddQ()

		// 現在の問題を保存
		currentProblem = Problem{
			Shape1:     shape1,
			Shape2:     shape2,
			correctSum:	correctSum,
			option: 	option,
			WithQ:      withQ,
		}

		c.JSON(http.StatusOK, gin.H{
			"shape1":     shape1,
			"shape2":     shape2,
			"option": 	  option,
			"withQ":      withQ,
		})
	})

	// 解答受け取り
	engine.POST("/answer", func(c *gin.Context) {
		var requestBody struct {
			Answer string `json:"answer"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "無効なリクエストです。"})
			return
		}

		// 解答を判定する
		correct, incorrect := judgement(requestBody.Answer, currentProblem.WithQ, currentProblem.option, currentProblem.correctSum)

		// 判定結果を表示
		fmt.Printf("ユーザーが選択した答え: %s\n", requestBody.Answer)
		fmt.Printf("正解数: %d, 不正解数: %d\n", correct, incorrect)

		// 判定結果を返す
		c.JSON(http.StatusOK, gin.H{"message": "答えを受け取りました。判定が完了しました。"})
	})

	// ランダムシードの設定
	rand.Seed(time.Now().UnixNano())
	engine.Run(":3000")
}