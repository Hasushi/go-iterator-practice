# 08: Collectors: Collect / Take / First

## 要件
次のユーティリティを実装して、既存の `Count` や `Filter` と組み合わせて使ってください。

- `Collect(seq iter.Seq[T]) []T` : シーケンスをスライスに溜める
- `Take(seq iter.Seq[T], n int) iter.Seq[T]` : 先頭 n 個だけ流す
- `First(seq iter.Seq[T]) (T, bool)` : 先頭 1 個だけ取り出す

テスト内容：
- 1..100 のうち偶数だけ取り出して先頭 5 個を `Collect` し、表示。
- 偶数の最初の 1 個を `First` で取得し、値と `ok` を表示。

## 実行
```bash
go run .
```
