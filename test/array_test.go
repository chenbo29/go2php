package test

import (
	"github.com/chenbo29/go2php/array"
	"testing"
)

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
