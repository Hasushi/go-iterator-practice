# 04: push→pull 変換（`iter.Pull`）

## 要件
- 1..∞（理論上）の自然数を生む `Naturals()` を `iter.Seq[int]` で実装。
- `iter.Pull` で pull 型に変換し、`next()` を 10 回だけ呼び出して表示する。
- 終了時には `defer stop()` を必ず呼び出す（約束事）。

## 実行
```bash
go run .
```
