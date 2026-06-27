package main

import "fmt"

func print_value() {
	// 1. 短縮宣言（型推論）
	name := "Gopher" // string型
	age := 10        // int型
	isTarget := true // bool型（真偽値）

	// 2. varを使った宣言（初期値なし = ゼロ値が入る）
	var score int // 0 が入る
	// var text string // ""（空文字）が入る

	// 3. 定数（一度決めたら変えられない値）
	const Pi = 3.14

	// 出力してみる
	fmt.Println("名前:", name)
	fmt.Println("年齢:", age)
	fmt.Println("ターゲットか:", isTarget)
	fmt.Println("初期値のスコア:", score)
	fmt.Println("円周率:", Pi)

	// 値の書き換え（:= は最初だけ、2回目以降は = を使う）
	age = 11
	fmt.Println("誕生日の後の年齢:", age)
}