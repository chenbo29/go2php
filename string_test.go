package go2php

import (
	"github.com/chenbo29/go2php/phpString"
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

func TestSubstr(t *testing.T) {
	str := "test"
	start := int64(10)
	result, err := phpString.Substr(str, start, 0)
	if err != nil {
		t.Logf(`Substr("%v", "%b", "%b") = %q ,want match for %q`, str, start, 0, result, err)
	}

	for i, test := range testsSubstr {
		result, err := phpString.Substr(test.str, test.start, test.length)
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

func TestTrim(t *testing.T) {
	for i, str := range testsTrim {
		ret := phpString.Trim(str.str, str.mask)
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

func TestLtrim(t *testing.T) {
	for i, str := range testsLtrim {
		ret := phpString.Ltrim(str.str, str.mask)
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

func TestRtrim(t *testing.T) {
	for i, str := range testsRtrim {
		ret := phpString.Rtrim(str.str, str.mask)
		if ret != str.result {
			t.Fatalf(`#%d Ltrim("%v", "%v") = %q, want result is %q`, i, str.str, str.mask, ret, str.result)
		}
	}
}
