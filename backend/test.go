package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
func generateOptions(correctSum int) (int) {
	rnd, option := random()
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
func judgement(input string, withQ bool, option int, correctSum int) (int,int) {
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
	} else if !withQ && input == "e"{
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

func provideProblem(){
	correctAnswernum := 0
	falseAnswernum := 0

	for i := 1; i <= 10; i++ {
		// 問題を生成
		shape1, shape2, correctSum := generateProblem()
		option := generateOptions(correctSum)

		questionPrefix, withQ := shouldAddQ()

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

		// 正解数と不正解数を更新
		correct, incorrect := judgement(input, withQ, option, correctSum)
		correctAnswernum += correct
		falseAnswernum += incorrect
	}

	fmt.Printf("\n10問が終了しました。お疲れ様でした！\n正解数: %d, 不正解数: %d\n", correctAnswernum, falseAnswernum)
}

func main() {

	//エンジン作成
	engine:=gin.Default()

	//CORSの許可
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

	//問題提供
	engine.GET("/problem", func(c *gin.Context){
		shape1, shape2, correctSum := generateProblem()
		_, withQ := shouldAddQ()

		c.JSON(http.StatusOK, gin.H{
			"shape1": shape1,
			"shape2": shape2,
			"correctSum": correctSum,
			"withQ": withQ,
		})
	})

	//解答受け取り
	engine.POST("/answer", func(c *gin.Context) {
		var requestBody struct {
			Answer string `json:"answer"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "無効なリクエストです。"})
			return
		}

		// 受け取った答えを処理する
		fmt.Printf("ユーザーが選択した答え: %s\n", requestBody.Answer)

		// 例として、正解メッセージを返す
		c.JSON(http.StatusOK, gin.H{"message": "答えを受け取りました。"})
	})


	rand.Seed(time.Now().UnixNano())
	engine.Run(":3000")
	//provideProblem()
}
