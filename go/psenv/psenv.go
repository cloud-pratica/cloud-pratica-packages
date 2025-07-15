package s3env

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
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

// Load Parameter Storeのパスを指定して、環境変数として設定する
func Load(ctx context.Context, path string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := ssm.NewFromConfig(cfg)
	res, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return errors.Wrap(err, "client.GetParameter")
	}

	value := res.Parameter.Value
	if value == nil {
		return errors.New("parameter value is nil")
	}
	envVars := parseEnvContent(*value)
	for key, value := range envVars {
		os.Setenv(key, value)
	}
	return nil
}
