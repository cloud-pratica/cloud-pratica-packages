# s3env

Amazon S3 に保存された`.env`ファイルから環境変数を読み込む Go パッケージです。

## 概要

`s3env`は、S3 バケットに保存された`.env`ファイルから環境変数を読み込むためのシンプルな方法を提供します。ローカルの`.env`ファイルを管理する必要がないアプリケーションに特に有用です。

## 機能

- S3 バケットから直接`.env`ファイルを読み込み
- 標準的な`.env`ファイル形式の解析
- コメント（`#`で始まる行）のサポート
- 環境変数の自動設定
- AWS SDK v2 との統合

## インストール

```bash
go get github.com/cloud-pratica/cloud-pratica-backend/packages/go/s3env
```

## 使用方法

### 基本的な使用方法

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/cloud-pratica/cloud-pratica-backend/packages/go/s3env"
)

func main() {
    ctx := context.Background()

    // S3から環境変数を読み込み
    err := s3env.Load(ctx, "my-bucket", "config/.env")
    if err != nil {
        log.Fatal(err)
    }

    // 環境変数が利用可能になります
    dbHost := os.Getenv("DB_HOST")
    apiKey := os.Getenv("API_KEY")
}
```

### 前提条件

1. **AWS 認証情報**: アプリケーションに適切な AWS 認証情報が設定されていることを確認してください。以下の方法で設定できます：

   - AWS CLI 設定（`aws configure`）
   - 環境変数（`AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`）
   - IAM ロール（EC2 インスタンスや ECS タスク用）
   - AWS SDK のデフォルト認証情報チェーン

2. **S3 権限**: AWS 認証情報に指定された S3 バケットとキーからの読み取り権限が必要です。

## .env ファイル形式

このパッケージは標準的な`.env`ファイル形式をサポートしています：

```env
# コメントがサポートされています
DB_HOST=localhost
DB_PORT=5432
API_KEY=your-api-key-here

# 空行は無視されます
DATABASE_URL=postgresql://user:pass@localhost:5432/db
```

### サポートされている機能

- **キー・バリューペア**: `KEY=value`形式
- **コメント**: `#`で始まる行は無視されます
- **空行**: 空白行はスキップされます
- **空白文字**: キーとバリューの前後の空白文字は削除されます

## API リファレンス

### `Load(ctx context.Context, bucket, key string) error`

S3 オブジェクトから環境変数を読み込み、現在のプロセス環境に設定します。

**パラメータ:**

- `ctx`: 操作のコンテキスト
- `bucket`: S3 バケット名
- `key`: S3 オブジェクトキー（.env ファイルへのパス）

**戻り値:**

- `error`: 操作が失敗した場合にエラーを返します

**例:**

```go
err := s3env.Load(ctx, "my-config-bucket", "environments/production/.env")
if err != nil {
    return fmt.Errorf("環境変数の読み込みに失敗しました: %w", err)
}
```

## エラーハンドリング

このパッケージは以下の場合にエラーを返します：

- AWS 設定の読み込みに失敗
- S3 オブジェクトの取得に失敗
- 無効な S3 バケットまたはキー

## 依存関係

- `github.com/aws/aws-sdk-go-v2/config` - AWS SDK v2 設定
- `github.com/aws/aws-sdk-go-v2/service/s3` - S3 サービスクライアント
- `github.com/aws/aws-sdk-go/aws` - AWS SDK v1 ユーティリティ
- `github.com/pkg/errors` - エラーラッピング

## テスト

以下のコマンドでテストを実行できます：

```bash
go test
```

テストスイートには、様々な`.env`ファイル形式とエッジケースが含まれており、堅牢な解析を保証します。

## ライセンス

このパッケージは cloud-pratica プロジェクトの一部です。
