# algorithm-golang

Golangを使用して基本的なソートアルゴリズムを学習・実装するためのリポジトリです。

## 📚 目的

このリポジトリは、以下の目的で作成されています：

- Golangでの基本的なソートアルゴリズムの実装を学ぶ
- アルゴリズムのパフォーマンスと計算量の理解
- 実行時間の計測と比較
- コーディングスキルの向上

## 🗂️ プロジェクト構造

```
algorithm-golang/
├── src/
│   ├── bogo/              # ボゴソート実装
│   ├── bubble/            # バブルソート実装
│   ├── cocktail/          # カクテルソート実装
│   └── pkg/
│       ├── executeTime.go          # 実行時間計測ユーティリティ
│       └── generateRandIntArray.go # ランダム配列生成ユーティリティ
├── go.mod
└── README.md
```

## 🚀 使い方

### 前提条件

- Go 1.21以上

### インストール

```bash
git clone https://github.com/jugeeem/algorithm-golang.git
cd algorithm-golang
```

### プログラムの実行

各アルゴリズムのディレクトリに移動して実行します：

```bash
# ボゴソートの実行
cd src/bogo
go run main.go

# バブルソートの実行
cd src/bubble
go run main.go

# カクテルソートの実行
cd src/cocktail
go run main.go
```

## 📖 実装済み・実装予定のアルゴリズム

### ソートアルゴリズム
- [x] ボゴソート (Bogo Sort)
- [x] バブルソート (Bubble Sort)
- [x] カクテルソート (Cocktail Sort)
- [ ] 選択ソート (Selection Sort)
- [ ] 挿入ソート (Insertion Sort)
- [ ] マージソート (Merge Sort)
- [ ] クイックソート (Quick Sort)
- [ ] ヒープソート (Heap Sort)

### ユーティリティ
- [x] 実行時間計測
- [x] ランダム整数配列生成

## 🔧 ユーティリティの使い方

### executeTime.go
アルゴリズムの実行時間を計測するためのユーティリティです。

```go
import "algorithm-golang/src/pkg"

// 同期実行（従来の方法）
pkg.Measure("Algorithm Name", func() {
    // アルゴリズム実行
})

// 非同期実行（goroutineを使用）
pkg.MeasureWithWait("Algorithm Name", func() {
    // アルゴリズム実行
})

// 複数のアルゴリズムを並行実行
pkg.MeasureConcurrent(map[string]func(){
    "Algorithm1": func() { /* 実行 */ },
    "Algorithm2": func() { /* 実行 */ },
})
```

### generateRandIntArray.go
ランダムな整数配列を生成するためのユーティリティです。

```go
import "algorithm-golang/src/pkg"

// 使用例（デフォルト: 100要素、範囲0-99999）
arr := pkg.GenerateRandIntArray()
```

## 👤 作成者

[jugeeem](https://github.com/jugeeem)
