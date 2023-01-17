package array

import "strings"

const CaseLower = "lower"
const CaseUpper = "upper"

// ChangeKeyCase array_change_key_case https://www.php.net/manual/en/function.array-change-key-case.php
func ChangeKeyCase[T comparable](arr map[string]T, t string) map[string]T {
	arrReturn := make(map[string]T)
	switch t {
	case CaseLower:
		for i, v := range arr {
			arrReturn[strings.ToLower(i)] = v
		}
	case CaseUpper:
		for i, v := range arr {
			arrReturn[strings.ToUpper(i)] = v
		}
	default:
		arrReturn = arr
	}
	return arrReturn
}

func Push[T comparable](arr []T, v T) (arrReturn []T) {
	arrReturn = append(arr, v)
	return
}

func Merge[T comparable](arrA []T, arrB []T) []T {
	for _, v := range arrB {
		arrA = append(arrA, v)
	}
	return arrA
}

func InArray[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
