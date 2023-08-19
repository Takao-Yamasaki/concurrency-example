package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	// 乱数ジェネレータの初期化（毎回異なる乱数を生成したいため）
	rand.Seed(time.Now().Unix())
	// 0から2999マイクロ秒までの間、ランダムにプログラムの停止
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Microsecond)

	num := rand.Intn(10)

	// ラッキーナンバーをチャネルに送信
	c <- num
}

func main() {
	fmt.Println("what is today's lucky number?")

	// チャネルを作成
	c := make(chan int)
	// チャネルを引数に渡す
	go getLuckyNum(c)

	// チャネルからラッキーナンバーを受信
	num := <-c

	fmt.Printf("Today's your luchey number is %d\n", num)

	// 使い終わったチャネルはcloseする
	close(c)
}
