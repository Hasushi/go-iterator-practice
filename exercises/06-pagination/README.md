# 06: ページネーション API をストリーム化

## 要件
- モック API `FetchPage(cursor string) (items []int, next string)` を用意（下記雛形あり）。
- これを包む `StreamAll()` を `iter.Seq[int]` で実装。内部で次ページを自動で辿る。
- ループ側は `for v := range StreamAll() { ... }` で値を受け取り、
  **20件受け取ったら `break`** して、以降のページ取得が走らないことを確認するためのログを出す。

## 実行
```bash
go run .
```
