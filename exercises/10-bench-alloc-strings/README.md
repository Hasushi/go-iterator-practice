# 10: ベンチ — 文字列の確保量を比較（iterator vs eager）

## 目的
大量の文字列を扱う処理で、
- **streaming（iterator）** と
- **一括保持（eagerなスライス）**  
の **allocs/op** と **B/op** の差を観察する。

## タスク
1. 10万〜50万件程度の「擬似行」を生成し、`len(s) > 8` のものだけ長さを合計する処理を、
   - Eager（`[]string` に保持 → 走査）
   - Iterator（`iter.Seq[string]`で逐次処理）
   で実装する。
2. ベンチを走らせ、`-benchmem` の値を比較。

## 実行例
```bash
go test -bench . -benchmem -run '^$' -count=3
```
