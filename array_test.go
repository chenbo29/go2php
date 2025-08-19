package go2php

import (
	"log"
	"strings"
	"testing"
)

func TestArrayAll(t *testing.T) {
	// 示例1：检查切片中所有元素是否为偶数
	numbers := []int{2, 4, 6, 8}
	allEven := ArrayAll(numbers, func(key, value interface{}) bool {
		num, ok := value.(int) // 类型断言，确保是int
		return ok && num%2 == 0
	})
	if !allEven {
		t.Fatalf("ArrayAll failed")
	}

	numbers = []int{2, 4, 6, 8, 9}
	allEven = ArrayAll(numbers, func(key, value interface{}) bool {
		num, ok := value.(int) // 类型断言，确保是int
		return ok && num%2 == 0
	})
	if allEven {
		t.Fatalf("ArrayAll failed")
	}

	// 示例2：检查map中所有动物名称长度是否小于12
	animals := map[string]string{
		"a": "dog",
		"b": "cat",
		"c": "cow",
		"d": "duck",
	}
	allShortNames := ArrayAll(animals, func(key, value interface{}) bool {
		name, ok := value.(string) // 类型断言，确保是string
		return ok && len(name) < 12
	})
	if !allShortNames {
		t.Fatalf("ArrayAll failed")
	}
	//超过6
	animals = map[string]string{
		"a": "dog",
		"b": "cat",
		"c": "cow",
		"d": "elephant",
	}
	allShortNames = ArrayAll(animals, func(key, value interface{}) bool {
		name, ok := value.(string) // 类型断言，确保是string
		return ok && len(name) < 6
	})
	if allShortNames {
		t.Fatalf("ArrayAll failed")
	}

	// 示例3：检查空集合-slice（返回true）
	var emptySlice []string
	emptyResult := ArrayAll(emptySlice, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if !emptyResult {
		t.Fatalf("ArrayAll failed")
	}
	//检查空集合-map
	var emptyMap = make(map[string]string)
	emptyMapResult := ArrayAll(emptyMap, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if !emptyMapResult {
		t.Fatalf("ArrayAll failed")
	}
	//不支持的类型（非切片/数组/ map）返回false
	var notSupport int64
	notSupportResult := ArrayAll(notSupport, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if notSupportResult {
		t.Fatalf("ArrayAll failed")
	}
}

func TestArrayAny(t *testing.T) {
	// 示例1：检查切片中任意一个元素是否为偶数
	numbers := []int{2, 4, 6, 8, 9}
	allEven := ArrayAny(numbers, func(key, value interface{}) bool {
		num, ok := value.(int) // 类型断言，确保是int
		return ok && num%2 == 0
	})
	if !allEven {
		t.Fatalf("ArrayAny failed")
	}

	numbers = []int{3, 9}
	allEven = ArrayAny(numbers, func(key, value interface{}) bool {
		num, ok := value.(int) // 类型断言，确保是int
		return ok && num%2 == 0
	})
	if allEven {
		t.Fatalf("ArrayAny failed")
	}

	// 示例2：检查map中所有动物名称长度是否小于12
	animals := map[string]string{
		"a": "dog",
		"b": "cat",
		"c": "cow",
		"d": "duck",
		"e": "duck-duck-duck",
	}
	allShortNames := ArrayAny(animals, func(key, value interface{}) bool {
		name, ok := value.(string) // 类型断言，确保是string
		return ok && len(name) < 12
	})
	if !allShortNames {
		t.Fatalf("ArrayAny failed")
	}
	//超过6
	animals = map[string]string{
		"a": "dog-dog",
		"b": "cat-cat",
		"c": "cow-cow",
		"d": "elephant",
	}
	allShortNames = ArrayAny(animals, func(key, value interface{}) bool {
		name, ok := value.(string) // 类型断言，确保是string
		return ok && len(name) < 6
	})
	if allShortNames {
		t.Fatalf("ArrayAny failed")
	}

	// 示例3：检查空集合-slice（返回true）
	var emptySlice []string
	emptyResult := ArrayAny(emptySlice, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if !emptyResult {
		t.Fatalf("ArrayAny failed")
	}
	//检查空集合-map
	var emptyMap = make(map[string]string)
	emptyMapResult := ArrayAny(emptyMap, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if !emptyMapResult {
		t.Fatalf("ArrayAny failed")
	}
	//不支持的类型（非切片/数组/ map）返回false
	var notSupport int64
	notSupportResult := ArrayAny(notSupport, func(key, value interface{}) bool {
		return false // 回调返回false，但空集合仍返回true
	})
	if notSupportResult {
		t.Fatalf("ArrayAny failed")
	}
}

func TestChangeKeyCase(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic error: %v", r)
		}
	}()

	test := map[string]int{"a": 1, "B": 2}

	retA := ArrayChangeKeyCase(test, CaseLower)
	var num int
	for i := range retA {
		if strings.ToLower(i) != i {
			t.Fatalf("ChangeKeyCase key upper error[%v][%v][%v]", test, num, i)
		}
		num++
	}
	num = 0
	retB := ArrayChangeKeyCase(test, CaseUpper)
	for i := range retB {
		if strings.ToUpper(i) != i {
			t.Fatalf("ChangeKeyCase key upper error[%v][%v][%v]", test, num, i)
		}
		num++
	}
	num = 0
	retUnknow := ArrayChangeKeyCase(test, "unknow")
	for i := range retUnknow {
		if !(i == "a" || i == "B") {
			t.Fatalf("ChangeKeyCase key “default” error[%v][%v][%v]", test, num, i)
		}
		num++
	}
}

func TestChunk(t *testing.T) {
	testA := []int{1, 2, 3, 4}
	retA := ArrayChunk(testA, 2)
	if len(retA) != 2 {
		t.Fatalf("Chunk error retA")
	}
	for i, v := range retA {
		for ii, vv := range v {
			if vv != i*2+ii+1 {
				t.Fatalf("Chunk error retB")
			}
		}
	}

	retB := ArrayChunk(testA, len(testA))
	if len(retB) != 1 {
		t.Fatalf("Chunk error retB")
	}
	for _, v := range retB {
		for i, vv := range v {
			if vv != i+1 {
				t.Fatalf("Chunk error retB")
			}
		}
	}

	retC := ArrayChunk(testA, len(testA)+1)
	if len(retC) != 1 {
		t.Fatalf("Chunk error retC")
	}
	for _, v := range retC {
		for i, vv := range v {
			if vv != i+1 {
				t.Fatalf("Chunk error retC")
			}
		}
	}
}

func TestArrayColumn(t *testing.T) {
	data := []map[string]int{
		{"a": 1, "b": 2, "c": 3, "d": 4},
		{"a": 5, "b": 6, "c": 7, "d": 8},
		{"a": 9, "b": 10, "c": 11, "d": 12},
	}
	ret := ArrayColumn(data, "b")
	for _, v := range ret {
		if ArrayInArray(v, []int{2, 6, 10}) == false {
			t.Fatalf("Column error")
		}
	}
}

func TestCombine(t *testing.T) {
	dataA := []string{"a", "b", "c", "d", "e", "f"}
	dataB := []string{"A", "B", "C", "D", "E", "F"}
	dataD := []string{"A", "B", "C", "D", "E"}
	dataC, err := ArrayCombine(dataA, dataB)
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range dataA {
		if dataC[v] != dataB[i] {
			t.Fatalf("Combine error")
		}
	}
	_, err = ArrayCombine(dataA, dataD)
	if err == nil {
		t.Fatalf("Combine Need error")
	}
}

func TestCountValues(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 2, 4, 1}
	ret := ArrayCountValues(data)
	if len(ret) != 5 {
		t.Fatalf("CountValues error")
	}
	if ret["1"] != 2 {
		t.Fatalf("CountValues error")
	}
}

func TestInArray(t *testing.T) {
	testA := []int{1, 2, 3}
	ArrayInArray(2, testA)
	for _, v := range testA {
		if v == 2 {
			return
		}
	}
	t.Fatalf("InArray error")
}

func TestDiff(t *testing.T) {
	dataA := []int{1, 2, 6, 4, 5}
	dataB := []int{1, 2, 3, 2, 4, 1}
	dataC := ArrayDiff(dataA, dataB)
	if ArrayInArray(6, dataC) == false || ArrayInArray(5, dataC) == false {
		log.Fatalf("Diff error")
	}
}

func TestDiffAssoc(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := map[string]int{"b": 2, "c": 4}
	c := map[string]int{"c": 3}

	out := ArrayDiffAssoc(a, b, c) // 只会保留 a 中不存在相同（key+value）的项
	if val, ok := out["a"]; ok && val == 1 {
		// success
	} else {
		t.Fatalf("InArray error")
	}
}

func TestArrayDiffUassoc(t *testing.T) {
	array1 := map[interface{}]interface{}{
		"a": "green",
		"b": "brown",
		"c": "blue",
		0:   "red", // 对应PHP中自动分配的索引键0
	}

	array2 := map[interface{}]interface{}{
		"a": "green",
		0:   "yellow", // 对应PHP中"yellow"的索引键0
		1:   "red",    // 对应PHP中"red"的索引键1
	}

	// 定义键比较函数（模拟PHP的<=>操作符）
	keyCompare := func(a, b interface{}) int {
		// 处理字符串键比较
		if aStr, okA := a.(string); okA {
			if bStr, okB := b.(string); okB {
				return strings.Compare(aStr, bStr)
			}
		}

		// 处理整数键比较
		if aInt, okA := a.(int); okA {
			if bInt, okB := b.(int); okB {
				if aInt < bInt {
					return -1
				} else if aInt > bInt {
					return 1
				}
				return 0
			}
		}

		// 不同类型的键视为不相等（例如字符串vs整数）
		return 1
	}

	// 计算差集
	result := ArrayDiffUassoc(array1, []map[interface{}]interface{}{array2}, keyCompare)

	for k, v := range result {
		if k == "b" && v != "brown" {
			log.Println(result)
			log.Fatalf("ArrayDiffUassoc error of key b")
		}
		if k == "c" && v != "blue" {
			log.Println(result)
			log.Fatalf("ArrayDiffUassoc error of key c")
		}
		if k == "0" && v != "red" {
			log.Println(result)
			log.Fatalf("ArrayDiffUassoc error of key 0")
		}

	}
}

func TestKeys(t *testing.T) {
	data := map[string]string{
		"a": "b",
		"c": "d",
		"e": "f",
	}
	ret := ArrayKeys(data)
	if len(ret) != 3 {
		t.Fatalf("Keys error")
	}
	if ArrayInArray("c", ret) == false {
		t.Fatalf("Keys error")
	}
}

func TestPush(t *testing.T) {
	testA := []int{1, 2, 3}
	testA = ArrayPush(testA, 4)
	ArrayInArray(4, testA)
}

func TestMerge(t *testing.T) {
	testA := []int{1, 2, 3}
	testB := []int{4, 5, 6}
	test := ArrayMerge(testA, testB)
	for i := range test {
		if i != test[i]-1 {
			t.Fatalf("TestMerge error: %v %v", i, test[i]-1)
		}
	}
	if len(test) != len(testA)+len(testB) {
		t.Fatalf("TestMerge error")
	}
}
