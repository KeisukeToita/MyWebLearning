# プロジェクト名

# 🗺️ 【完全統合版】Web技術スタック制覇マトリクス

| 開発順 | 対象スタック / レイヤー | 想定工数 | 難易度 | 業界での「主戦場」 | 【合格クリア条件】（これが書けたら卒業） | 推奨テスト / デプロイ先 | リンク(シート) |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **0-1** | Docker<br>(インフラ共通基盤) | 5h | ★☆☆☆☆ | 開発環境の標準規格 | `docker compose up -d` でDB起動・初期データ投入が完全に自動完了する。 | WSL2環境 | [リンク](https://drive.google.com/file/d/1bS2M4ZCfzeChi9yrhl6dNF-3USxJlW5m/view?usp=drive_link) |
| **0-2** | Next.js<br>(フロント共通基盤) | 20h | ★★☆☆☆ | モダンWebのデファクト | API通信部を `repository` 層に隠蔽し、環境変数で接続先を切り替えられる。 | Vercel | - |
| **0-3** | JWT (JSON Web Token)<br>(認証共通基盤) | 3h | ★★☆☆☆ | ステートレス認証の標準 | バックエンド側でユーザーDBを参照せず、トークンの検証（デコード）のみで認証を完結させる。 | 共通仕様として固定 | - |
| **1** | Go × Echo<br>(API主導) | 10h | ★★☆☆☆ | 高負荷マイクロサービス | `net/http/pprof` を仕込み、負荷テスト時のメモリ/CPUプロファイルを可視化する。 | testing + Testcontainers<br>→ Fly.io | - |
| **2** | Rust × Axum<br>(API主導) | 25h | ★★★★★ | Web3 / 超低レイテンシAPI | リンター `clippy` の警告をゼロにし、ORM（`sqlx`）のコンパイル時SQL型チェックを通す。 | cargo test<br>→ Shuttle | - |
| **3** | Java × Spring Boot<br>(API主導) | 15h | ★★★☆☆ | 金融・大規模Enterprise | コントローラーを手書きせず、OpenAPI Generatorによるスキーマ駆動で実装する。 | JUnit 5 + Mockito<br>→ Render | - |
| **4** | C# × ASP.NET Core<br>(API主導) | 12h | ★★★☆☆ | ゲーム会社 / 外資系Enterprise | EF Coreでクエリを書く際、変更追跡をオフにする `AsNoTracking()` を明示的に使い分けて高速化する。 | xUnit<br>→ Azure (無料枠) | - |
| **5** | Python × FastAPI<br>(API主導) | 10h | ★★☆☆☆ | AI / データ分析API基盤 | `async def` 内で同期的な重い処理（AI推論等）を `run_in_threadpool` に逃がす非同期の基本を書く。 | Pytest<br>→ Render | - |
| **6** | Ruby × Ruby on Rails<br>(モノリス) | 8h | ★☆☆☆☆ | スタートアップ / 最速開発 | Gem `bullet` を導入し、N+1クエリを完全にゼロにする。Hotwireでの非同期更新を1箇所組む。 | RSpec<br>→ Render | - |
| **7** | PHP × Laravel<br>(モノリス) | 12h | ★★☆☆☆ | 中規模自社サービス / 受託 | `FormRequest` を独立させ、バリデーションロジックをコントローラーから完全に追い出す。 | PHPUnit (Pest)<br>→ Render | - |
| **8** | Python × Django<br>(モノリス) | 10h | ★★☆☆☆ | 社内業務・データ管理インフラ | 標準の管理画面をカスタムし、**「選択したデータをCSVで一括DLするアクション」**を生やす。 | Django標準 TestCase<br>→ Render | - |
| **9** | Elixir × Phoenix<br>(モノリス) | 20h | ★★★★☆ | 大規模チャット / リアルタイム | Phoenix PubSubを使い、「片方の画面の操作が、別ブラウザへ0.1秒で同期する」機能をJS無しで組む。 | ExUnit<br>→ Fly.io | - |
| **10** | TS × Nuxt 3<br>(フロント派生) | 12h | ★★☆☆☆ | 国内モダンフロント主軸 | `server/api/` を使い、**外部APIキーをブラウザに一切露出させないBFF（Backend For Frontend）**を組む。 | Vitest<br>→ Vercel | - |
| **11** | TS × SvelteKit<br>(フロント派生) | 10h | ★★☆☆☆ | 海外モダンWeb / 軽量メディア | Form Actions を用い、**「ブラウザのJSを無効化してもフォーム送信が動く」**設計にする。 | Playwright<br>→ Cloudflare | - |
| **12** | TS × Hono<br>(フロント派生) | 10h | ★★★☆☆ | エッジサーバーレス / マイクロ | Node.js固有API（`fs`等）を排し、純粋な Web Standard API の記述のみで Cloudflare Workers 上で動かす。 | Vitest<br>→ Cloudflare | - |
