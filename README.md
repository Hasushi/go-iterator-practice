# Go イテレータ練習問題（Go 1.23）

Go 1.23 の `for range` over function と標準パッケージ `iter` を使って、
実践的に学べる練習問題セットです。**すべて解答（`solutions/`）付き**。

## 前提
- Go 1.23 以降（`iter` パッケージが使えること）

## 使い方
- 問題は `exercises/` ディレクトリ、解答は `solutions/` ディレクトリにあります。
- まずは `exercises/<番号-名前>/` に入り、`go run .` が通るように TODO を埋めてください。
- 解答を確認したい場合は `solutions/<番号-名前>/` を `go run .` してください。

例：
```bash
cd solutions/01-fib-basic
go run .
```

---

## 目次
1. 01-fib-basic: フィボナッチの基本イテレータ  
2. 02-early-break: `break` での早期停止と片付け  
3. 03-map-filter: Map/Filter アダプタの実装  
4. 04-pull-conversion: push→pull 変換（`iter.Pull`）  
5. 05-file-walk: ファイル走査をイテレータ化し中断可能に  
6. 06-pagination: ページネーション API をストリーム化  
7. 07-error-propagation: エラー伝搬パターン（`Seq2[T,error]`）  
8. 08-collectors: `Collect`/`Take`/`First` ユーティリティ

各ディレクトリの `README.md` に詳細な要件が書かれています。
