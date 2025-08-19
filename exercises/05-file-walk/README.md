# 05: ファイル走査をイテレータ化

## 要件
- `WalkFiles(root string) iter.Seq[string]` を実装し、ファイルパスを1つずつ `yield`。
- ディレクトリはスキップ。
- ループ側で「最初に見つけた `.go` ファイルで `break`」できるようにする。

## 実行
任意のルートを指定：
```bash
go run . .
```
