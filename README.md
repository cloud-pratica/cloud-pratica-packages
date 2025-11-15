# cloud-pratica-packages

cloud-pratica プロジェクトで使用する共通パッケージのリポジトリです。

## このリポジトリについて

このリポジトリは、複数の Go パッケージと Protocol Buffers 定義を管理するモノレポです。各パッケージは独立してバージョン管理され、Git タグを使用してリリースされます。

## ディレクトリ構造

```
cloud-pratica-packages/
├── go/                    # Go パッケージ
│   ├── errs/             # エラーハンドリング
│   ├── grpcerrs/         # gRPC エラー
│   ├── logging/          # ロギング
│   ├── psenv/            # Parameter Store 環境変数
│   └── s3env/            # S3 環境変数
├── proto/                 # Protocol Buffers 定義
│   └── cost-provider/    # cost-provider サービスの proto 定義
└── Makefile              # 共通の Makefile
```

## パッケージを公開する手順

### 1. 現在のバージョンを確認

パッケージを公開する前に、現在のバージョンを確認します。

```bash
# パッケージのディレクトリに移動
cd go/errs

# 現在のバージョンを確認
make version
# 出力例: go/errs/v0.1.0
```

このコマンドは、現在のディレクトリの `version` ファイルからバージョン情報を読み取り、`<相対パス>/<バージョン>` の形式で表示します。

### 2. 変更をコミットしてリモートにプッシュ

コードの変更や `version` ファイルの更新がある場合は、まずコミットしてリモートリポジトリにプッシュします。

```bash
# 変更をステージング
git add .

# コミット（適切なコミットメッセージを付ける）
git commit -m "Update package to v0.1.1"

# リモートリポジトリにプッシュ
git push origin main
# または
git push origin <ブランチ名>
```

**注意**: `make push` コマンドはタグのみをプッシュします。コードの変更をリモートに反映させるには、この手順でコミットをプッシュする必要があります。

### 3. 新しいバージョンでタグを作成してプッシュ

新しいバージョンを指定して Git タグを作成し、リモートリポジトリにプッシュします。

```bash
# 新しいバージョンを指定してプッシュ
make push VERSION=v0.1.1
```

**重要**: `make push` を実行すると、以下の処理が行われます：

1. 指定されたバージョンで Git タグが作成されます（形式: `<相対パス>/<バージョン>`）
   - 例: `go/errs/v0.2.0`
2. タグがリモートリポジトリにプッシュされます
3. **`version` ファイルが新しいバージョンで上書きされます**

### 4. バージョン形式について

バージョンは [Semantic Versioning](https://semver.org/) の形式に従う必要があります：

- 形式: `v<メジャー>.<マイナー>.<パッチ>`
- 例: `v1.0.0`, `v0.2.1`, `v2.1.3`

`make push` コマンドは、指定されたバージョンが正しい形式かどうかを自動的にチェックします。

## 補足

### version ファイルの自動更新

`make push` コマンドを実行すると、**現在のディレクトリの `version` ファイルが自動的に新しいバージョンで上書きされます**。

例：

```bash
# 現在の version ファイルの内容
cat go/errs/version
# v0.1.0

# 新しいバージョンでプッシュ
make push VERSION=v0.2.0

# version ファイルが更新される
cat go/errs/version
# v0.2.0
```

この動作により、`version` ファイルは常に最新のリリースバージョンを保持します。

### 新規ディレクトリを作成する場合

新しいパッケージやディレクトリを作成する際は、**必ず `version` ファイルを準備してください**。

```bash
# 新しいパッケージディレクトリを作成
mkdir -p go/newpackage

# version ファイルを作成（初期バージョンを設定）
echo "v0.1.0" > go/newpackage/version

# 必要に応じて Makefile も作成（ルートの Makefile を include）
echo "include ../../Makefile" > go/newpackage/Makefile
```

`version` ファイルがない状態で `make push` を実行すると、エラーが発生します。

## 使用例

### 例 1: パッチバージョンの更新

```bash
cd go/errs

# 現在のバージョンを確認
make version
# go/errs/v0.1.0

# 変更をコミットしてプッシュ（必要な場合）
git add .
git commit -m "Fix bug in error handling"
git push origin main

# パッチバージョンを更新してタグをプッシュ
make push VERSION=v0.1.1
```

### 例 2: マイナーバージョンの更新

```bash
cd go/logging

# 現在のバージョンを確認
make version
# go/logging/v0.1.2

# 変更をコミットしてプッシュ（必要な場合）
git add .
git commit -m "Add new logging features"
git push origin main

# マイナーバージョンを更新してタグをプッシュ
make push VERSION=v0.2.0
```

### 例 3: メジャーバージョンの更新

```bash
cd go/psenv

# 現在のバージョンを確認
make version
# go/psenv/v0.1.5

# 変更をコミットしてプッシュ（必要な場合）
git add .
git commit -m "Breaking changes: refactor API"
git push origin main

# メジャーバージョンを更新してタグをプッシュ
make push VERSION=v1.0.0
```

## Makefile コマンド一覧

### `make version`

現在のディレクトリの `version` ファイルからバージョン情報を読み取り、`<相対パス>/<バージョン>` の形式で表示します。

```bash
make version
# 出力例: go/errs/v0.1.0
```

### `make push VERSION=<バージョン>`

指定されたバージョンで Git タグを作成し、リモートリポジトリにプッシュします。また、`version` ファイルを新しいバージョンで更新します。

```bash
make push VERSION=v1.0.0
```

**前提条件:**

- `VERSION` 引数が必須です
- `VERSION` は Semantic Versioning 形式（`v1.0.0`）である必要があります
- 現在のディレクトリに `version` ファイルが存在する必要があります

## パッケージの使用

各パッケージは独立してバージョン管理されているため、Go モジュールとして個別にインポートできます。

```go
import (
    "github.com/cloud-pratica/cloud-pratica-packages/go/errs"
    "github.com/cloud-pratica/cloud-pratica-packages/go/logging"
)
```

特定のバージョンを使用する場合は、`go get` コマンドでタグを指定します：

```bash
go get github.com/cloud-pratica/cloud-pratica-packages/go/errs@go/errs/v0.2.0
```

## コントリビューション

新しいパッケージを追加する場合は、以下の手順に従ってください。

1. 新しいディレクトリを作成
2. **必ず `version` ファイルを作成**（初期バージョンを設定、例: `v0.1.0`）
3. 必要に応じて `Makefile` を作成（ルートの Makefile を include）
4. パッケージのコードとドキュメントを追加
5. 初回リリース時は `make push VERSION=v0.1.0` を実行

## ライセンス

このリポジトリは cloud-pratica プロジェクトの一部です。
