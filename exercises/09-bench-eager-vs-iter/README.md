# 09: ベンチ比較 — イテレータ vs スライス vs チャネル

## 目的
以下の3実装の **速度（ns/op）** と **メモリ（B/op, allocs/op）** を比較し、
各アプローチのトレードオフを理解する。

- Eager（スライスで中間結果を保持）
- Iterator（`iter.Seq`で遅延パイプライン）
- Channel（goroutine+chanでストリーミング）

## タスク
1. `*_test.go` の TODO を埋めて、各実装を完成させる。
2. 既定の `N=1_000_000, K=50_000` でベンチを実行：
   ```bash
   go test -bench . -benchmem -run '^$' -count=3
   ```
3. `N` と `K` を変えながら傾向を観察し、レポート（所感）を `NOTES.md` にまとめる。
   - 例: `N=1e6, K=1e6` / `N=1e6, K=1e3`
   - 期待：allocs/op は **Iterator < Channel << Eager** になりやすい。

## 比較対象の処理内容
- 1..N を生成 → 偶数だけ残す（Filter） → 2乗する（Map） → 先頭 K 個を合計
