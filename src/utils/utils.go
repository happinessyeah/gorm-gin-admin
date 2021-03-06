package utils

import (
	"math/rand"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var (
	RegIDCard   = regexp.MustCompile("(^[1-9]\\d{5}\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}$)|(^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$)")
	RegPhoneNum = regexp.MustCompile("^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$")
	RegEmail    = regexp.MustCompile("^([\\w-_]+(?:\\.[\\w-_]+)*)@((?:[a-z0-9]+(?:-[a-zA-Z0-9]+)*)+\\.[a-z]{2,6})$")
)

// SQLColumnToHumpStyle sql转换成驼峰模式
func SQLColumnToHumpStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if i > 0 && in[i-1] == '_' && in[i] != '_' {
			s := strings.ToUpper(string(in[i]))
			ret += s
		} else if in[i] == '_' {
			continue
		} else {
			ret += string(in[i])
		}
	}
	return
}

// HumpToSQLColumnStyle 驼峰转sql
func HumpToSQLColumnStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if unicode.IsUpper(rune(in[i])) {
			ret += "_" + strings.ToLower(string(in[i]))
		} else {
			ret += string(in[i])
		}
	}
	return
}

// PubGetEnvString
func PubGetEnvString(key string, defaultValue string) (ret string) {
	ret = os.Getenv(key)
	if len(ret) == 0 {
		ret = defaultValue
	}
	return
}

// PubGetEnvBool
func PubGetEnvBool(key string, defaultValue bool) (ret bool) {
	val := strings.ToLower(os.Getenv(key))
	if val == "true" {
		ret = true
	} else if val == "false" {
		ret = false
	} else {
		ret = defaultValue
	}
	return
}

func SpiltInterfaceList(d string, sep string) []interface{} {
	var (
		list []interface{}
	)
	for _, v := range strings.Split(d, sep) {
		list = append(list, v)
	}
	return list
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// RandStringBytesMask
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return "id_" + string(b)
}
