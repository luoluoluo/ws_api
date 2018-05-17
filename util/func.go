package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// string to int
func ParseInt(v string) int {
	if s, err := strconv.Atoi(v); err == nil {
		return s
	}
	return 0
}

// str to float
func ParseFloat(v string) float64 {
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		return s
	}
	return 0
}

// str to bool
func ParseBool(v string) bool {
	if s, err := strconv.ParseBool(v); err == nil {
		return s
	}
	return false
}

// 生成32位随机id
func RandId() string {
	return Md5(strconv.FormatInt(time.Now().UnixNano(), 10) + RandCode(6, "string"))
}

// md5
func Md5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

//生成随机字符串
func RandCode(strlen int, codetype string) string {
	var codes string
	switch codetype {
	case "string":
		codes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	case "number":
		codes = "0123456789"
	default:
		return ""
	}
	codeLen := len(codes)
	data := make([]byte, strlen)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < strlen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
