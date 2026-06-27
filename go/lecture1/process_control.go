package main

import "fmt"

func processControl() {
	fmt.Println("--- 1. if のテスト ---")
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "は偶数です")
	} else {
		fmt.Println(num, "は奇数です")
	}

	fmt.Println("\n--- 2. for のテスト ---")
	// 1から5までの数字を足し算してみる
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
		fmt.Printf("%d を足しました（現在の合計: %d）\n", i, sum)
	}

	fmt.Println("\n--- 3. switch のテスト ---")
	day := "土曜日"
	switch day {
	case "土曜日", "日曜日":
		fmt.Println("お休みです！Goの勉強をしましょう。")
	default:
		fmt.Println("平日です。お仕事や学校を頑張りましょう。")
	}
}