package main

import "fmt"

func sliceAndMap() {
	fmt.Println("--- 1. スライスのテスト ---")
	// スライスの作成
	tasks := []string{"環境構築", "変数の学習", "制御構文の学習"}
	
	// appendで要素を追加
	tasks = append(tasks, "ポインタの学習")
	tasks = append(tasks, "スライスの学習")

	// for range で中身をすべて表示（インデックス番号は使わないので _ で捨てる）
	for _, task := range tasks {
		fmt.Println("完了したタスク:", task)
	}

	fmt.Println("\n--- 2. マップのテスト ---")
	// マップの作成（商品の価格表）
	prices := map[string]int{
		"コーヒー": 400,
		"紅茶":    450,
	}

	// 新しい商品を追加
	prices["ケーキ"] = 600

	// for range で中身を表示
	for item, price := range prices {
		fmt.Printf("%s は %d 円です\n", item, price)
	}

	fmt.Println("\n--- 3. マップの安全な値の取り出し ---")
	// マップから値を取り出すとき、2つ目の戻り値(ok)で「そのキーが存在するか」を確認できる
	target := "ジュース"
	val, ok := prices[target]
	if ok {
		fmt.Println(target, "の価格は", val, "円です")
	} else {
		fmt.Println(target, "はメニューにありません！") // ここが実行される
	}
}