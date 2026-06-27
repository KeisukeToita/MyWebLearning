package main

import "fmt"

// 1. 構造体の定義
type Player struct {
	Name  string
	Level int
}

// 2. 値レシーバのメソッド（コピーを受け取る）
// ※この中でLevelを変えても、元のPlayerには影響しない
func (p Player) ShowStatus() {
	fmt.Printf("【ステータス】名前: %s | レベル: %d\n", p.Name, p.Level)
}

// 3. ポインタレシーバのメソッド（住所を受け取る）
// ※元のPlayerのデータを直接書き換えることができる
func (p *Player) LevelUp() {
	p.Level++ // レベルを1上げる
	fmt.Println(p.Name, "はレベルアップした！")
}

// （参考）悪い例：値レシーバでレベルを上げようとするメソッド
func (p Player) BadLevelUp() {
	p.Level++
	fmt.Println("(内部) コピーのレベルを上げました")
}


func structAndPointerExample() {
	// 構造体の初期化
	hero := Player{
		Name:  "勇者Gopher",
		Level: 1,
	}

	hero.ShowStatus()

	fmt.Println("\n--- 悪い例（値レシーバ）を試す ---")
	hero.BadLevelUp()
	hero.ShowStatus() // レベルは1のまま！

	fmt.Println("\n--- 良い例（ポインタレシーバ）を試す ---")
	hero.LevelUp()
	hero.ShowStatus() // レベルが2になっている！
}