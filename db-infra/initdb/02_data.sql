-- テストユーザーの挿入
INSERT INTO users (name, email) VALUES 
('山田 太郎', 'yamada@example.com'),
('佐藤 美咲', 'sato@example.com'),
('鈴木 一郎', 'suzuki@example.com');

-- テストタスクの挿入（ユーザーID 1と2に紐付け）
INSERT INTO tasks (user_id, title, is_completed) VALUES 
(1, '設計書のレビューを行う', false),
(1, 'Docker環境の動作確認', true),
(2, 'フロントエンドの画面実装', false);