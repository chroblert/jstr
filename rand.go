package jstr

import (
	"math/rand"
	"time"
)

// 产生随机几个字符
func GenerateRandomString(length int) string {
	// 可选的字符集合
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机字符串
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
