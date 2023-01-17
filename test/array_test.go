package test

import (
	"github.com/chenbo29/go2php/array"
	"testing"
)

func TestChangeKeyCase(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic error: %v", r)
		}
	}()

	test := map[string]int{"a": 1, "B": 2}

	retA := array.ChangeKeyCase(test, array.CaseLower)
	var num int
	for i := range retA {
		if num == 0 && i != "a" {
			t.Fatalf("ChangeKeyCase key lower error")
		}
		if num == 1 && i != "b" {
			t.Fatalf("ChangeKeyCase key lower error")
		}
		num++
	}
	num = 0
	retB := array.ChangeKeyCase(test, array.CaseUpper)
	for i := range retB {
		if num == 0 && i != "A" {
			t.Fatalf("ChangeKeyCase key upper error")
		}
		if num == 1 && i != "B" {
			t.Fatalf("ChangeKeyCase key upper error")
		}
		num++
	}
}

func TestInArray(t *testing.T) {
	testA := []int{1, 2, 3}
	array.InArray(testA, 2)
	for _, v := range testA {
		if v == 2 {
			return
		}
	}
	t.Fatalf("InArray error")
}

func TestPush(t *testing.T) {
	testA := []int{1, 2, 3}
	testA = array.Push(testA, 4)
	array.InArray(testA, 4)
}

func TestMerge(t *testing.T) {
	testA := []int{1, 2, 3}
	testB := []int{4, 5, 6}
	test := array.Merge(testA, testB)
	for i := range test {
		if i != test[i]-1 {
			t.Fatalf("TestMerge error: %v %v", i, test[i]-1)
		}
	}
	if len(test) != len(testA)+len(testB) {
		t.Fatalf("TestMerge error")
	}
}
