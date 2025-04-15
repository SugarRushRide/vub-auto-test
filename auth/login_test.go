//  -
// Created: 2025/4/15

package auth

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	token, err := getToken("chengjun.jiang", "shanghai0303")
	if err != nil {
		t.Fatalf("获取 token 失败: %v", err)
	}
	if token == "" {
		t.Fatal("token 为空")
	}
	t.Logf("✅ 成功获取 token: %s", token)
}
