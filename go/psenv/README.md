# psenv

AWS Systems Manager Parameter Store に保存された `.env` ファイル形式のパラメータから環境変数を読み込む Go パッケージです。

## 概要

`psenv`は、AWS Systems Manager Parameter Store に保存された `.env` ファイル形式のパラメータから環境変数を読み込むためのシンプルな方法を提供します。機密情報や設定値を Parameter Store で管理し、アプリケーションで安全に利用するのに特に有用です。

## 機能

- Parameter Store から直接 `.env` ファイル形式のパラメータを読み込み
- 標準的な `.env` ファイル形式の解析
- コメント（`#`で始まる行）のサポート
- 環境変数の自動設定
- AWS SDK v2 との統合
- 暗号化されたパラメータの自動復号化

## インストール

```bash
go get github.com/cloud-pratica/cloud-pratica-packages/go/psenv
```

## 使用方法

### 基本的な使用方法

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/cloud-pratica/cloud-pratica-packages/go/psenv"
)

func main() {
    ctx := context.Background()

    // Parameter Store から環境変数を読み込み
    err := psenv.Load(ctx, "/my-app/production/.env")
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

2. **Parameter Store 権限**: AWS 認証情報に指定された Parameter Store パラメータからの読み取り権限が必要です。

## Parameter Store での設定

Parameter Store に `.env` ファイル形式のパラメータを保存する例：

```bash
# AWS CLI を使用してパラメータを保存
aws ssm put-parameter \
    --name "/my-app/production/.env" \
    --value "DB_HOST=localhost
DB_PORT=5432
API_KEY=your-api-key-here
DATABASE_URL=postgresql://user:pass@localhost:5432/db" \
    --type "SecureString" \
    --description "Production environment variables"
```

### サポートされている形式

Parameter Store に保存する値は標準的な `.env` ファイル形式である必要があります：

```env
# コメントがサポートされています
DB_HOST=localhost
DB_PORT=5432
API_KEY=your-api-key-here

# 空行は無視されます
DATABASE_URL=postgresql://user:pass@localhost:5432/db
```

### サポートされている機能

- **キー・バリューペア**: `KEY=value` 形式
- **コメント**: `#` で始まる行は無視されます
- **空行**: 空白行はスキップされます
- **空白文字**: キーとバリューの前後の空白文字は削除されます
- **暗号化**: SecureString タイプのパラメータは自動的に復号化されます

## API リファレンス

### `Load(ctx context.Context, path string) error`

Parameter Store のパラメータから環境変数を読み込み、現在のプロセス環境に設定します。

**パラメータ:**

- `ctx`: 操作のコンテキスト
- `path`: Parameter Store のパラメータパス（例：`/my-app/production/.env`）

**戻り値:**

- `error`: 操作が失敗した場合にエラーを返します

**例:**

```go
err := psenv.Load(ctx, "/my-app/production/.env")
if err != nil {
    return fmt.Errorf("環境変数の読み込みに失敗しました: %w", err)
}
```

## エラーハンドリング

このパッケージは以下の場合にエラーを返します。

- AWS 設定の読み込みに失敗
- Parameter Store パラメータの取得に失敗
- 無効なパラメータパス
- パラメータ値が nil の場合

## 依存関係

- `github.com/aws/aws-sdk-go-v2/config` - AWS SDK v2 設定
- `github.com/aws/aws-sdk-go-v2/service/ssm` - Systems Manager サービスクライアント
- `github.com/aws/aws-sdk-go/aws` - AWS SDK v1 ユーティリティ
- `github.com/pkg/errors` - エラーラッピング

## セキュリティ

- Parameter Store の SecureString タイプを使用することで、機密情報を暗号化して保存できます
- このパッケージは自動的に暗号化されたパラメータを復号化します
- IAM ロールを使用することで、認証情報を安全に管理できます

## ライセンス

このパッケージは cloud-pratica プロジェクトの一部です。
