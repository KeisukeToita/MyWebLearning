package main

import (
	"fmt"
	"time"
)

// 裏で実行させる関数（引数にチャネルを受け取る）
func worker(id int, ch chan string) {
	fmt.Printf("ワーカー%d: 重い作業を開始します...\n", id)

	// time.Sleepで2秒間待機（重い処理をシミュレーション）
	time.Sleep(2 * time.Second)

	fmt.Printf("ワーカー%d: 作業完了！\n", id)

	// チャネルに結果の文字列を送信
	ch <- fmt.Sprintf("【ワーカー%dの計算結果】", id)
}

func goRoutineExample() {
	fmt.Println("--- メイン処理 開始 ---")

	// 1. 文字列を通す土管（チャネル）を作る
	resultCh := make(chan string)

	// 2. ゴルーチンを起動！（裏で worker 関数が走り始める）
	go worker(1, resultCh)
	go worker(2, resultCh)

	fmt.Println("メイン処理: ワーカーの終了を待っています...")

	// 3. チャネルからデータを受信する
	// ※ワーカーがチャネルにデータを入れるまで、メイン処理はここでストップして待つ
	result1 := <-resultCh
	result2 := <-resultCh

	fmt.Println("メイン処理: 結果を受け取りました！ ->", result1)
	go worker(3, resultCh)
	result3 := <-resultCh
	fmt.Println("メイン処理: 結果を受け取りました！ ->", result2)
	fmt.Println("メイン処理: 結果を受け取りました！ ->", result3)

	fmt.Println("--- メイン処理 終了 ---")
}
