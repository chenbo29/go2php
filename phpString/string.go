package phpString

import (
	"errors"
	"strings"
)

func Substr(str string, start int64, length int64) (string, error) {
	strRune := []rune(str)
	if start > int64(len(str)) {
		err := errors.New("start长度不能大于字符串长度")
		return "", err
	}
	if start < 0 {
		start = start + int64(len(strRune))
	}
	if length < 0 {
		length = length + int64(len(strRune))
	}
	if length == 0 {
		length = int64(len(strRune))
	}
	return string([]rune(str)[start : start+length]), nil
}

func Trim(str string, characterMask string) string {
	return strings.Trim(str, characterMask)
}

func Ltrim(str string) string {
	return strings.TrimLeft(str, "")
}

func Rtrim(str string) string {
	return strings.TrimRight(str, "")
}
