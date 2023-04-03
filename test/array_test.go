package test

import (
	"github.com/chenbo29/go2php/array"
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

	retA := array.ChangeKeyCase(test, array.CaseLower)
	var num int
	for i := range retA {
		if strings.ToLower(i) != i {
			t.Fatalf("ChangeKeyCase key upper error[%v][%v][%v]", test, num, i)
		}
		num++
	}
	num = 0
	retB := array.ChangeKeyCase(test, array.CaseUpper)
	for i := range retB {
		if strings.ToUpper(i) != i {
			t.Fatalf("ChangeKeyCase key upper error[%v][%v][%v]", test, num, i)
		}
		num++
	}
}

func TestChunk(t *testing.T) {
	testA := []int{1, 2, 3, 4}
	retA := array.Chunk(testA, 2)
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

	retB := array.Chunk(testA, len(testA))
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

	retC := array.Chunk(testA, len(testA)+1)
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
	data := []map[string]any{
		{"a": 1, "b": 2, "c": 3, "d": 4},
		{"a": 5, "b": 6, "c": 7, "d": 8},
		{"a": 9, "b": 10, "c": 11, "d": 12},
	}
	ret := array.Column(data, "b")
	for _, v := range ret {
		if array.InArray(v, []any{2, 6, 10}) == false {
			t.Fatalf("Column error")
		}
	}
}

func TestCombine(t *testing.T) {
	dataA := []string{"a", "b", "c", "d", "e", "f"}
	dataB := []string{"A", "B", "C", "D", "E", "F"}
	dataC := array.Combine(dataA, dataB)
	for i, v := range dataA {
		if dataC[v] != dataB[i] {
			t.Fatalf("Combine error")
		}
	}
}

func TestCountValues(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 2, 4, 1}
	ret := array.CountValues(data)
	if len(ret) != 5 {
		t.Fatalf("CountValues error")
	}
	if ret["1"] != 2 {
		t.Fatalf("CountValues error")
	}
}

func TestInArray(t *testing.T) {
	testA := []int{1, 2, 3}
	array.InArray(2, testA)
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
	array.InArray(4, testA)
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
