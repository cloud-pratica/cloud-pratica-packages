# Protocol Buffers (proto) ディレクトリ

このディレクトリは、gRPC サービスで使用する Protocol Buffers（プロトコルバッファー）の定義ファイルと、そこから生成された Go コードを管理するためのディレクトリです。

## 📚 このディレクトリについて

### Protocol Buffers とは？

Protocol Buffers（通称：protobuf、proto）は、Google が開発したデータシリアライゼーション形式です。gRPC サービス間の通信で使用されるデータ構造や API の定義を記述するために使われます。

**簡単に言うと：**

- `.proto`ファイル：サービスやデータ構造の「設計図」
- 生成された`.pb.go`ファイル：その設計図から自動生成された Go 言語のコード

### このディレクトリの役割

- **API 定義の管理**: gRPC サービスのインターフェースを`.proto`ファイルで定義
- **コード生成**: `.proto`ファイルから Go 言語のコードを自動生成
- **共有ライブラリ**: 他のプロジェクトからこの定義をインポートして使用

## 📁 ディレクトリ構造

```
proto/
├── cost-provider/
│   ├── api/
│   │   └── v1/
│   │       └── service.proto          # ⭐ 編集OK: API定義ファイル
│   └── gen/
│       ├── service.pb.go              # ❌ 編集禁止: 自動生成されたコード
│       └── service_grpc.pb.go         # ❌ 編集禁止: 自動生成された gRPC コード
├── go.mod                             # Go の依存関係管理
├── go.sum                             # Go の依存関係のチェックサム
├── Makefile                           # 生成コマンドをまとめたファイル
└── README.md                          # このファイル
```

## ⚠️ 重要な注意事項

### ❌ 編集してはいけないファイル

以下のファイルは**絶対に手動で編集しないでください**。これらは`.proto`ファイルから自動生成されるため、編集しても次回の生成時に上書きされてしまいます。編集したい場合は、必ず `api/` 配下の `.proto` を更新してから再生成します。

- `cost-provider/gen/*.pb.go` - すべての生成された Go / gRPC コード

### ✅ 編集して良いファイル

- `cost-provider/api/v1/service.proto` - API 定義ファイル（ここを編集して変更を加えます）

## 🛠️ MacBook での環境構築

### 1. Homebrew のインストール（未インストールの場合）

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 2. Protocol Buffer Compiler (protoc) のインストール

```bash
brew install protobuf
```

インストールが完了したら、バージョンを確認します：

```bash
protoc --version
# 例: libprotoc 5.29.3
```

### 3. Go 言語のプラグインのインストール

Protocol Buffers から Go コードを生成するために、以下の 2 つのプラグインが必要です：

```bash
# Protocol Buffers用のGoプラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# gRPC用のGoプラグイン
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 4. PATH の設定

インストールしたプラグインが使用できるように、`~/.zshrc`（または`~/.bash_profile`）に以下を追加します：

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

設定を反映させます：

```bash
source ~/.zshrc
```

### 5. 動作確認

以下のコマンドで、プラグインが正しくインストールされているか確認できます：

```bash
protoc-gen-go --version
protoc-gen-go-grpc --version
```

## 🔨 proto ファイルから Go コードを生成する手順

### 基本的な生成コマンド

このディレクトリ（`proto/`）で以下のいずれかの方法で生成します。

### 1. Makefile を使う（推奨）

```bash
make generate-proto
# 別サービスを生成したい場合の例
make generate-proto SERVICE=another-service
```

`Makefile` の `SERVICE` 変数を切り替えることで、同じ構造の別サービスにも対応できます（デフォルトは `cost-provider`）。

### 2. 手動で `protoc` を実行する

```bash
protoc \
  --proto_path=cost-provider/api/v1 \
  --go_out=cost-provider/gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=cost-provider/gen \
  --go-grpc_opt=paths=source_relative \
  service.proto
```

### コマンドの説明

- `--proto_path=cost-provider/api/v1`: `.proto` ファイルを探すベースパス
- `--go_out=cost-provider/gen`: Go コードの出力先
- `--go_opt=paths=source_relative`: 相対パスで出力
- `--go-grpc_out=cost-provider/gen`: gRPC 用 Go コードの出力先
- `--go-grpc_opt=paths=source_relative`: gRPC コードも相対パスで出力
- `service.proto`: 入力となる proto ファイル（`--proto_path` で指定したディレクトリから解決されます）

### 生成後の確認

コマンド実行後、以下のファイルが生成（または更新）されます：

- `cost-provider/gen/service.pb.go`
- `cost-provider/gen/service_grpc.pb.go`

## 📝 proto ファイルの編集方法

### 基本的な流れ

1. **proto ファイルを編集**: `cost-provider/api/v1/service.proto`を編集
2. **コードを生成**: 上記の`protoc`コマンドを実行
3. **生成されたコードを確認**: `gen/`ディレクトリ内のファイルが更新される

### 例：新しいメッセージを追加する場合

`service.proto`に以下のように追加：

```protobuf
message NewMessage {
  string field1 = 1;
  int32 field2 = 2;
}
```

その後、`protoc`コマンドを実行すると、`service.pb.go`に`NewMessage`構造体が自動生成されます。

## 🔍 現在の API 定義の概要

現在定義されているサービス：

- **CostProviderService**: クラウドプロバイダーのコスト情報を取得するサービス
  - `GetCost`: 指定期間のコスト情報を取得

サポートされているプロバイダー：

- AWS
- GCP
- Datadog

## 🚀 他のプロジェクトでの使用方法

この proto 定義を他の Go プロジェクトで使用する場合：

```go
import (
    costproviderpb "github.com/cloud-pratica/cloud-pratica-packages/proto/cost-provider/gen"
)
```

## ❓ よくある質問

### Q: proto ファイルを編集した後、必ずコード生成が必要ですか？

A: はい。proto ファイルを編集したら、必ず`protoc`コマンドを実行して Go コードを再生成してください。生成しないと、変更が反映されません。

### Q: 生成されたコードにエラーが出る場合は？

A: まず、proto ファイルの構文エラーを確認してください。`protoc`コマンド実行時にエラーメッセージが表示されます。

### Q: 複数の proto ファイルがある場合は？

A: 各 proto ファイルに対して`protoc`コマンドを実行するか、ワイルドカードを使用できます：

```bash
protoc --proto_path=cost-provider/api/v1 \
       --go_out=cost-provider/gen --go_opt=paths=source_relative \
       --go-grpc_out=cost-provider/gen --go-grpc_opt=paths=source_relative \
       *.proto
```

## 📚 参考資料

- [Protocol Buffers 公式ドキュメント](https://protobuf.dev/)
- [gRPC 公式ドキュメント](https://grpc.io/)
- [Go Protocol Buffers ガイド](https://protobuf.dev/getting-started/gotutorial/)

## 🔗 関連リンク

- [Protocol Buffers 言語ガイド](https://protobuf.dev/programming-guides/proto3/)
- [gRPC Go クイックスタート](https://grpc.io/docs/languages/go/quickstart/)
