package main

import (
	"errors"
	"fmt"
)

// 1. 基本的な関数
func greet(name string) string {
	return "こんにちは、" + name + "さん！"
}

// 2. 割り算をする関数（エラーを返す可能性がある）
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New() で新しいエラーメッセージを作る
		return 0, errors.New("エラー: ゼロで割ることはできません（0除算）")
	}
	return a / b, nil
}

func funcAndErrorHandling() {
	fmt.Println("--- 1. 関数の呼び出し ---")
	message := greet("Gopher")
	fmt.Println(message)

	fmt.Println("\n--- 2. エラーハンドリング（成功パターン） ---")
	// 10 ÷ 2 を計算する
	result, err := divide(10.0, 2.0)
	
	// err が nil じゃない（＝エラーが発生した）かをチェックする
	if err != nil {
		fmt.Println(err) // エラーメッセージを表示
	} else {
		fmt.Printf("計算成功: 答えは %v です\n", result)
	}

	fmt.Println("\n--- 3. エラーハンドリング（失敗パターン） ---")
	// 10 ÷ 0 を計算しようとしてみる
	result2, err2 := divide(10.0, 0.0)
	
	if err2 != nil {
		fmt.Println(err2) // ここを通るはず！
	} else {
		fmt.Printf("計算成功: 答えは %v です\n", result2)
	}
}