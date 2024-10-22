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
type Session struct{
	CorrectCount	int
	WrongCount		int
	CurrentCount	int
	Finished		bool
}

var currentProblem Problem
var session Session

// ランダムな問題を生成
func generateProblem() (string, string, int, bool) {
	shapes := []string{"〇", "△", "□", "☆"} // 図形のリスト

	// ランダムに2つの図形を選択
	shape1 := shapes[rand.Intn(len(shapes))]
	shape2 := shapes[rand.Intn(len(shapes))]

	// 合計を計算
	sum := shapeValues[shape1] + shapeValues[shape2]
	withQ := rand.Float64() < 0.1
	return shape1, shape2, sum, withQ
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

// 正解の判定
func judgement(input string, withQ bool, option int, correctSum int) (bool) {
	AnswerJudge := false

	if withQ && input == "q" {
		fmt.Println("正解。!です。")
		AnswerJudge = true
	} else if !withQ && input == "w" {
		if option == correctSum {
			fmt.Println("正解。", correctSum, "です。")
			AnswerJudge = true
		} else {
			fmt.Println("不正解です。正解は", correctSum, "です。")
		}
	} else if !withQ && input == "e" {
		if option != correctSum {
			fmt.Println("正解。Eです。")
			AnswerJudge = true
		} else {
			fmt.Println("不正解です。正解は", correctSum, "です。")
		}
	} else {
		fmt.Println("不正解です。あなたの入力は", input, "です。")
	}
	return AnswerJudge
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
		shape1, shape2, correctSum, withQ := generateProblem()
		option :=generateOptions(correctSum)

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

		// JSONのバインドに失敗した場合のエラーハンドリング
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "無効なリクエストです。"})
			return
		}

		// 解答を判定する
		correct := judgement(requestBody.Answer, currentProblem.WithQ, currentProblem.option, currentProblem.correctSum)

		// 正解・不正解のカウント
		if correct {
			session.CorrectCount++
		} else {
			session.WrongCount++
		}
		session.CurrentCount++

		// 10回回答されたら終了状態へ移行するための処理
		if session.CurrentCount == 10 {
			session.Finished = true
		}
		// レスポンスを返す
		message := "正解です。"
		if !correct {
			message = "不正解です。"
		}

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	// 集計結果を取得するエンドポイント
	engine.GET("/summary", func(c *gin.Context) {
		if session.Finished {
			c.JSON(http.StatusOK, gin.H{
				"correctCount": session.CorrectCount,
				"wrongCount":   session.WrongCount,
			})
			session = Session{} // セッションをリセット
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "まだ10問回答されていません。"})
		}
	})

	// ランダムシードの設定
	rand.Seed(time.Now().UnixNano())
	engine.Run(":3000")
}