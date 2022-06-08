package phpString

import (
	"errors"
	"strings"
)

// Substr https://www.php.net/manual/zh/function.substr.php
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

// Trim https://www.php.net/manual/zh/function.trim.php
// todo 使用..，可以指定字符的范围。
func Trim(str string, characterMask string) string {
	if characterMask == "" {
		defaultCMask := []string{" ", "\t", "\n", "\r", "\x0b"}
		for _, mask := range defaultCMask {
			str = strings.Trim(str, mask)
		}
		return str
	} else {
		var cMM []string
		Permutation(&cMM, characterMask, 0)
		for _, cm := range cMM {
			temp := []rune(cm)
			for _, t := range temp {
				str = strings.Trim(str, string(t))
			}
		}
		return str
	}
}

// Ltrim https://www.php.net/manual/zh/function.ltrim.php
func Ltrim(str string, characterMask string) string {
	if characterMask == "" {
		defaultCMask := []string{" ", "\t", "\n", "\r", "\x0b"}
		for _, mask := range defaultCMask {
			str = strings.TrimLeft(str, mask)
		}
		return str
	} else {
		var cMM []string
		Permutation(&cMM, characterMask, 0)
		for _, cm := range cMM {
			temp := []rune(cm)
			for _, t := range temp {
				str = strings.TrimLeft(str, string(t))
			}
		}
		return str
	}
}

// Rtrim https://www.php.net/manual/zh/function.rtrim.php
func Rtrim(str string, characterMask string) string {
	if characterMask == "" {
		defaultCMask := []string{" ", "\t", "\n", "\r", "\x0b"}
		for _, mask := range defaultCMask {
			str = strings.TrimRight(str, mask)
		}
		return str
	} else {
		var cMM []string
		Permutation(&cMM, characterMask, 0)
		for _, cm := range cMM {
			temp := []rune(cm)
			for _, t := range temp {
				str = strings.TrimRight(str, string(t))
			}
		}
		return str
	}
}

// Chr https://www.php.net/manual/zh/function.chr.php
func Chr(num int) (string, error) {
	if num > 31 && num < 127 {
		return string(rune(num)), nil
	} else {
		return "", errors.New("暂不支持控制字符【第 0~31 个字符以及第 127 个字符】")
	}
}
