package utils

import (
	"errors"
	"math/rand"
	"os"
	"time"
)

func RemoveDuplicatesString(arr []string) []string {
	encountered := make(map[string]bool)
	result := make([]string, 0)

	for v := range arr {
		if !encountered[arr[v]] {
			encountered[arr[v]] = true
			result = append(result, arr[v])
		}
	}

	return result
}

func GetRandomString(randomLen int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, randomLen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < randomLen; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}

	return string(result)
}

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if !os.IsNotExist(err) {
			return true
		}
		return false
	}
	return true
}

func TypeDirectToInt(val interface{}) (ret int64, err error) {
	switch val.(type) {
	case uint:
		return int64(val.(uint)), nil
	case uint8:
		return int64(val.(uint8)), nil
	case uint16:
		return int64(val.(uint16)), nil
	case uint32:
		return int64(val.(uint32)), nil
	case uint64:
		return int64(val.(uint64)), nil
	case int:
		return int64(val.(int)), nil
	case int8:
		return int64(val.(int8)), nil
	case int16:
		return int64(val.(int16)), nil
	case int32:
		return int64(val.(int32)), nil
	case int64:
		return int64(val.(int64)), nil
	case float32:
		return int64(val.(float32)), nil
	case float64:
		return int64(val.(float64)), nil
	}

	return 0, errors.New("return error int")
}
