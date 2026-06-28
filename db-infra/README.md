## 🐳 開発環境（DB基盤）の起動方法

### 1. 起動（初回データ投入も自動で行われます）
$ docker compose up -d

### 2. 動作確認用のデータ確認コマンド
$ docker compose exec db psql -U user -d my_database -c "SELECT * FROM users;"

### 3. 完全に初期状態（データ全削除）に戻してやり直したい時
$ docker compose down -v