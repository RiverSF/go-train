package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func CreateUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func String(ad interface{}) string {
	b, err := json.Marshal(ad)
	if err != nil {
		return fmt.Sprintf("%v", ad)
	}
	return string(b)
}

func HmacSha256(ad string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(ad))
	return hex.EncodeToString(h.Sum(nil))
}

const NullMd5 = "d41d8cd98f00b204e9800998ecf8427e"

func Md5(str string) string {
	if len(str) == 0 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func StrIsNum(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func StringToFloat64(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return res
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%v", value)
}

func StringToInt(value string) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return res
}

func StringToInt64(value string) int64 {
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func Int64ToString(int64 int64) string {
	return fmt.Sprintf("%v", int64)
}

func IsContainHanStr(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func HanStrToUnicode(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}

func StringFirstToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.ToLower(str)
	first := str[0:1]
	other := str[1:]

	return strings.ToUpper(first) + other
}
