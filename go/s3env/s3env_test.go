package s3env

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseEnvContent(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     map[string]string
	}{
		{
			name:     "case1",
			fileName: "case1.env",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
				"KEY3": "value3",
			},
		},
		{
			name:     "case2",
			fileName: "case2.env",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:     "case3",
			fileName: "case3.env",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:     "case4",
			fileName: "case4.env",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:     "case5",
			fileName: "case5.env",
			want: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
		},
		{
			name:     "case6",
			fileName: "case6.env",
			want:     map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := os.ReadFile(fmt.Sprintf("testdata/%s", tt.fileName))
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			got := parseEnvContent(string(content))

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected result (-got +want):\n%s", diff)
			}
		})
	}
}
