package s3env

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// parseEnvContent .envファイルの内容を解析して環境変数のマップを返す
func parseEnvContent(content string) map[string]string {
	envVars := make(map[string]string)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		envVars[key] = value
	}

	return envVars
}

// Load S3の.envファイルを読み込んで、環境変数として設定する
func Load(ctx context.Context, bucket, key string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	out, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("client.GetObject: %w", err)
	}
	defer out.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(out.Body)

	envVars := parseEnvContent(buf.String())
	for key, value := range envVars {
		os.Setenv(key, value)
	}
	return nil
}
