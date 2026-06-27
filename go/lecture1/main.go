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
	// learning Go with DevContainer
	var message string = "Hello, Go World from DevContainer!"
	fmt.Println(message)
	print_value()
	processControl()
	funcAndErrorHandling()
	structAndPointerExample()
	sliceAndMap()
	goRoutineExample()
}
