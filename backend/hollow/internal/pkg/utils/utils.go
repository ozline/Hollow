package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateTokenSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetTimestamp13() int64 {
	return time.Now().UnixNano() / 1e6
}

func GenerateMD5Random() string {
	h := md5.New()
	h.Write([]byte(Int64ToBytes(GetTimestamp13())))
	return hex.EncodeToString(h.Sum(nil))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// 生成6位数字验证码
func GetCapature() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// interface 转 string
// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value := value.(type) {
	case float64:
		ft := value
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value
		key = strconv.Itoa(it)
	case uint:
		it := value
		key = strconv.Itoa(int(it))
	case int8:
		it := value
		key = strconv.Itoa(int(it))
	case uint8:
		it := value
		key = strconv.Itoa(int(it))
	case int16:
		it := value
		key = strconv.Itoa(int(it))
	case uint16:
		it := value
		key = strconv.Itoa(int(it))
	case int32:
		it := value
		key = strconv.Itoa(int(it))
	case uint32:
		it := value
		key = strconv.Itoa(int(it))
	case int64:
		it := value
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value
		key = strconv.FormatUint(it, 10)
	case string:
		key = value
	case []byte:
		key = string(value)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func Intval(t1 interface{}) int {
	var t2 int
	switch t1 := t1.(type) {
	case uint:
		t2 = int(t1)
	case int8:
		t2 = int(t1)
	case uint8:
		t2 = int(t1)
	case int16:
		t2 = int(t1)
	case uint16:
		t2 = int(t1)
	case int32:
		t2 = int(t1)
	case uint32:
		t2 = int(t1)
	case int64:
		t2 = int(t1)
	case uint64:
		t2 = int(t1)
	case float32:
		t2 = int(t1)
	case float64:
		t2 = int(t1)
	case string:
		t2, _ = strconv.Atoi(t1)
	default:
		t2 = t1.(int)
	}
	return t2
}
