package array

import (
	"fmt"
	"log"
	"strings"
)

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

func Chunk[T comparable](arr []T, length int) (arrReturn [][]T) {
	start := 0
	end := length
	for end <= len(arr) {
		arrReturn = append(arrReturn, arr[start:end])
		start += length
		end += length
	}
	if len(arrReturn)*length < len(arr) {
		arrReturn = append(arrReturn, arr[end-length:])
	}
	return
}

func Column[T comparable](arr []map[string]T, key string) (arrReturn []T) {
	for _, v := range arr {
		if _, ok := v[key]; ok {
			arrReturn = append(arrReturn, v[key])
		}
	}
	return
}

func Combine[T comparable](arrA []T, arrB []T) map[T]T {
	if len(arrA) != len(arrB) {
		log.Fatalln("Combine length mismatch")
	}
	arrReturn := make(map[T]T, len(arrA))
	for i := range arrA {
		arrReturn[arrA[i]] = arrB[i]
	}
	return arrReturn
}

func CountValues[T comparable](arr []T) map[string]int {
	arrReturn := make(map[string]int, len(arr))
	for _, v := range arr {
		arrReturn[fmt.Sprintf("%#v", v)]++
	}
	return arrReturn
}

func Diff[T comparable](arrA []T, arrB []T) (arrReturn []T) {
	for _, v := range arrA {
		if InArray(v, arrB) == false {
			arrReturn = append(arrReturn, v)
		}
	}
	return
}

func Keys[T comparable](arr map[string]T) (arrReturn []string) {
	for k := range arr {
		arrReturn = append(arrReturn, k)
	}
	return
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

func InArray[T comparable](elem T, arr []T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
