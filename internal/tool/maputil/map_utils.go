package maputil

import (
	"fmt"
	"strconv"
)

func CollectKeysOfMap[K int | int64 | string, V any](m map[K]V) []K {
	if len(m) == 0 {
		return []K{}
	}
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func GetInt64(mp map[string]any, key string) (int64, bool) {
	if len(mp) == 0 {
		return 0, false
	}
	val := mp[key]
	if val == nil {
		return 0, false
	}
	str := fmt.Sprint(val)
	if len(str) == 0 {
		return 0, false
	}
	parsed, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parse %s to int64 error: %v\n", str, err)
		return 0, false
	}
	return parsed, true
}

func GetBool(mp map[string]any, key string) (val bool, ok bool) {
	if len(mp) == 0 {
		return false, false
	}
	v := mp[key]
	if v == nil {
		return false, false
	}
	str := fmt.Sprint(v)
	if len(str) == 0 {
		return false, false
	}
	parsed, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Printf("parse %s to bool error: %v\n", str, err)
		return false, false
	}
	return parsed, true
}

func GetString(mp map[string]any, key string) (string, bool) {
	if len(mp) == 0 {
		return "", false
	}
	val := mp[key]
	if val == nil {
		return "", false
	}
	return fmt.Sprint(val), true
}
