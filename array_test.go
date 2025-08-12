package go2php

import (
	"log"
	"strings"
	"testing"
)

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

func TestDiff(t *testing.T) {
	dataA := []int{1, 2, 6, 4, 5}
	dataB := []int{1, 2, 3, 2, 4, 1}
	dataC := ArrayDiff(dataA, dataB)
	if ArrayInArray(6, dataC) == false || ArrayInArray(5, dataC) == false {
		log.Fatalf("Diff error")
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
