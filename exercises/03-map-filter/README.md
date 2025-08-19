# 03: Map / Filter アダプタを実装

## 要件
- `Map(f, seq)` と `Filter(pred, seq)` を `iter.Seq` ベースで実装してください。
- 入力: 1..100 のシーケンス
  - 偶数のみを `Filter`
  - それを2乗する `Map`
  - 最初の 10 個だけ表示して `break`

**注意**: `Filter` 内部でも `yield=false` を伝播させて、即停止できるように。

## 実行
```bash
go run .
```
