# 07: エラー伝搬パターン（Seq2[T, error]）

## 要件
- 行入力を模したイテレータ `Lines(limit int)` を作成し、
  `iter.Seq2[string, error]` で `(line, err)` を yield してください。
- 想定：10 行に一度エラーを起こすようにし、エラーが出たら
  「エラーを最後に1回だけ流して終了」する実装にする（以降は値を生成しない）。
- ループ側は `for line, err := range Lines(100) { ... }` と書き、
  `err != nil` のときにメッセージを出して `break`。

## 実行
```bash
go run .
```
