package go2php

import (
	"errors"
	"fmt"
	"strings"
)

const CaseLower = "lower"
const CaseUpper = "upper"

// ArrayChangeKeyCase https://www.php.net/manual/en/functArrayion.array-change-key-case.php
func ArrayChangeKeyCase[T comparable](arr map[string]T, t string) map[string]T {
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

func ArrayChunk[T comparable](arr []T, length int) (arrReturn [][]T) {
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

func ArrayColumn[T comparable](arr []map[string]T, key string) (arrReturn []T) {
	for _, v := range arr {
		if _, ok := v[key]; ok {
			arrReturn = append(arrReturn, v[key])
		}
	}
	return
}

// ArrayCombine https://www.php.net/manual/zh/function.array-combine.php
func ArrayCombine[T comparable](arrA []T, arrB []T) (map[T]T, error) {
	if len(arrA) != len(arrB) {
		return nil, errors.New("combine array length mismatch")
	}
	arrReturn := make(map[T]T, len(arrA))
	for i := range arrA {
		arrReturn[arrA[i]] = arrB[i]
	}
	return arrReturn, nil
}

func ArrayCountValues[T comparable](arr []T) map[string]int {
	arrReturn := make(map[string]int, len(arr))
	for _, v := range arr {
		arrReturn[fmt.Sprintf("%#v", v)]++
	}
	return arrReturn
}

// ArrayDiffAssoc 返回 base 中那些在任何 others 中都找不到（key AND value 相同）的键值对。https://www.php.net/manual/zh/functArrayion.array-diff-assoc.php
func ArrayDiffAssoc[K comparable, V comparable](base map[K]V, others ...map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range base {
		found := false
		for _, other := range others {
			if ov, ok := other[k]; ok && ov == v {
				found = true
				break
			}
		}
		if !found {
			result[k] = v
		}
	}
	return result
}

func ArrayDiff[T comparable](arrA []T, arrB []T) (arrReturn []T) {
	for _, v := range arrA {
		if ArrayInArray(v, arrB) == false {
			arrReturn = append(arrReturn, v)
		}
	}
	return
}

func ArrayKeys[T comparable](arr map[string]T) (arrReturn []string) {
	for k := range arr {
		arrReturn = append(arrReturn, k)
	}
	return
}

func ArrayPush[T comparable](arr []T, v T) (arrReturn []T) {
	arrReturn = append(arr, v)
	return
}

func ArrayMerge[T comparable](arrA []T, arrB []T) []T {
	for _, v := range arrB {
		arrA = append(arrA, v)
	}
	return arrA
}

func ArrayInArray[T comparable](elem T, arr []T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
