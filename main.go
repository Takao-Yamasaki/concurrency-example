package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")

	// sync.WaitGroup構造体のwgを用意
	var wg sync.WaitGroup

	// wgの内部カウンタの値を+1
	wg.Add(1)

	go func() {
		// ゴルーチンが終了したときにwgの内部カウンタの値を-1するよう設定
		defer wg.Done()
		getLuckyNum()
	}()

	// 内部カウンタが0になるまでメインゴルーチンをブロックして待つ
	wg.Wait()
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
