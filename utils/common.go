package utils

import (
	"encoding/hex"
	"github.com/google/uuid"
)

// GetUUID32 生成一个不带横杠的 32 位 UUID 字符串
func GetUUID32() string {
	id := uuid.New()
	return hex.EncodeToString(id[:])
}

// MapKeys 接受任意类型的 map，并返回其所有键组成的切片
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
