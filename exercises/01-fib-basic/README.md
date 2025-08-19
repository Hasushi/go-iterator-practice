# 01: フィボナッチの基本イテレータ

## 要件
`Fibonacci(limit int) iter.Seq[int]` を実装してください。
- 0,1,1,2,3,5,... と増やし、値が `limit` を超えたら生成を止める。
- `for num := range Fibonacci(100)` で値を受け取り、20 を超えたら `break` して終了する。
- `break` によりイテレータ側で `yield=false` を検知し、メッセージを出して終了する。

## 実行
```bash
go run .
```
