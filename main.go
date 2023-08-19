package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")
	// 新ゴールーチン
	go getLuckyNum()

	// 5秒停止
	time.Sleep(time.Second * 5)
}

func getLuckyNum() {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	// 乱数ジェネレータの初期化（毎回異なる乱数を生成したいため）
	rand.Seed(time.Now().Unix())
	// 0から2999マイクロ秒までの間、ランダムにプログラムの停止
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Microsecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your luchey number is %d\n", num)
}
