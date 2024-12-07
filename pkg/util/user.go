package util

import (
	"fmt"
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

// 生成4位随机字母（包含大小写字母）
func generateRandomLetters(length int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 大小写字母
	rand.Seed(uint64(time.Now().UnixNano()))                               // 使用时间戳作为随机数种子

	var sb strings.Builder
	for i := 0; i < length; i++ {
		index := rand.Intn(len(letters)) // 随机选择一个字母
		sb.WriteByte(letters[index])     // 将字母添加到字符串中
	}
	return sb.String()
}

// GenerateName 生成带有4位字母和时间戳的字符串
func GenerateName() string {
	randomLetters := generateRandomLetters(4)            // 生成4个随机字母
	timestamp := time.Now().Format("20060102150405")     // 获取当前时间戳（精确到秒）
	return fmt.Sprintf("%s%s", randomLetters, timestamp) // 拼接字母和时间戳
}
