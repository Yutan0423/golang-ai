package util

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	res := os.Getenv(key)
	if res == "" {
		fmt.Printf("環境変数 %v が設定されていません", key)
	}
	return res
}
