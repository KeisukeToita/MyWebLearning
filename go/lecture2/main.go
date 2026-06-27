package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// レスポンス用のJSONの形を構造体で定義する
// ※右側の `json:"message"` は「JSONに変換するときは小文字の message にしてね」という指示です
type Response struct {
	Message string `json:"message"`
}

// リクエストが来たときに実行される関数（ハンドラー）
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. レスポンスのヘッダーを「JSONを返しますよ」という設定にする
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	// 2. 返したいデータ（構造体）を作成
	res := Response{
		Message: fmt.Sprintf("Hello, %s! Welcome to Go Web API World!", name),
	}

	// 3. 構造体をJSONデータに変換して、クライアントに送る（書き込む）
	json.NewEncoder(w).Encode(res)
}

func main() {

	// "/api/hello" というURLにリクエストが来たら、helloHandler を実行するように登録
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("サーバーを起動中... (http://localhost:8080)")

	// 8080番ポートでサーバーを起動して待ち受ける
	// ※エラーが発生した場合は log.Fatal などで落とすのが一般的ですが、まずはシンプルに
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}
}
