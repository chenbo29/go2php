package go2php

import (
	"errors"
	"regexp"
	"testing"
)

type Substr struct {
	str      string
	expected string
	start    int64
	length   int64
}

type Trim struct {
	str    string
	mask   string
	result string
}

var testsSubstr = []Substr{
	{"abcdef", "bcdef", 1, 0},
	{"abcdef", "bcd", 1, 3},
	{"abcdef", "cde", 2, -1},
	{"abcdef", "", 1, 3},
	{"abcdef", "abcd", 0, 4},
	{"abcdef", "abcde", 0, -1},
	{"abcdef", "de", -3, -1},
	{"abcdef", "ef", -2, 0},
	{"abcdef", "d", -3, 1},
}

func TestStringSubstr(t *testing.T) {
	str := "test"
	start := int64(10)
	result, err := StringSubstr(str, start, 0)
	if err == nil {
		t.Fatalf(`Substr("%v", "%b", "%b") = %q ,want match for %q`, str, start, 0, result, errors.New("start长度不能大于字符串长度"))
	}

	for i, test := range testsSubstr {
		result, err = StringSubstr(test.str, test.start, test.length)
		want := regexp.MustCompile(test.expected)
		if !want.MatchString(result) || err != nil {
			t.Fatalf(`#%d Substr("%v", "%b", "%b") = %q, %q ,want match for %q, nil`, i, test.str, test.start, test.length, result, err, want)
		}
	}
}

var testsTrim = []Trim{
	{"\t\tThese are a few words :) ...  ", "", "These are a few words :) ..."},
	{"\t\tThese are a few words :) ...  ", " \t.", "These are a few words :)"},
	{"Hello World", "Hdle", "o Wor"},
	{"欢迎光临", "临", "欢迎光"},
}

func TestStringTrim(t *testing.T) {
	for i, str := range testsTrim {
		ret := StringTrim(str.str, str.mask)
		if ret != str.result {
			t.Fatalf(`#%d Trim("%v", "%v") = %q, want result is %q`, i, str.str, str.mask, ret, str.result)
		}
	}
}

var testsLtrim = []Trim{
	{"\t\tThese are a few words :) ...  ", "", "These are a few words :) ...  "},
	{"\t\tThese are a few words :) ...  ", " \t.", "These are a few words :) ...  "},
	{"Hello World", "Hdle", "o World"},
}

func TestStringLtrim(t *testing.T) {
	for i, str := range testsLtrim {
		ret := StringLtrim(str.str, str.mask)
		if ret != str.result {
			t.Fatalf(`#%d Ltrim("%v", "%v") = %q, want result is %q`, i, str.str, str.mask, ret, str.result)
		}
	}
}

var testsRtrim = []Trim{
	{"\t\tThese are a few words :) ...  ", "", "\t\tThese are a few words :) ..."},
	{"\t\tThese are a few words :) ...  ", " \t.", "\t\tThese are a few words :)"},
	{"Hello World", "Hdle", "Hello Wor"},
}

func TestStringRtrim(t *testing.T) {
	for i, str := range testsRtrim {
		ret := StringRtrim(str.str, str.mask)
		if ret != str.result {
			t.Fatalf(`#%d Ltrim("%v", "%v") = %q, want result is %q`, i, str.str, str.mask, ret, str.result)
		}
	}
}

type Chr struct {
	target int
	expect string
}

var testsChr = []Chr{
	{target: 65, expect: "A"},
	{target: 97, expect: "a"},
	{target: 33, expect: "!"},
	{target: 63, expect: "?"},
	{target: 32, expect: " "},
}

func TestStringChr(t *testing.T) {
	for i, v := range testsChr {
		ret, err := StringChr(0)
		if err == nil {
			t.Fatalf(`#%d Chr("%d") = %q, want result is %q`, i, v.target, ret, errors.New("暂不支持控制字符【第 0~31 个字符以及第 127 个字符】"))
		}
		ret, err = StringChr(v.target)
		if err != nil {
			t.Fatalf(`#%d Chr("%d") = %q, want result is %q`, i, v.target, ret, err)
		}
		if ret != v.expect {
			t.Fatalf(`#%d Chr("%d") = %q, want result is %q`, i, v.target, ret, v.expect)
		}
	}
}

func TestStringInArrayString(t *testing.T) {
	if StringInArrayString("a", []string{"a", "b", "c"}) != true {
		t.Fatalf(`StringInArrayString("a", []string{"a", "b", "c"}) = false, want true`)
	}
	if StringInArrayString("a", []string{"b", "c"}) != false {
		t.Fatalf(`StringInArrayString("a", []string{"a", "b", "c"}) = true, want false`)
	}
}
