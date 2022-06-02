package go2php

import (
	"fmt"
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

func TestTrim(t *testing.T) {
	if errMsg := testTrim("\t\tThese are a few words ...  ", "These are a few words ...", ""); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testTrim("\t\tThese are a few words ...  ", "These are a few words ", "\t."); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testTrim("Hello World", "o Wor", "Hdle"); errMsg != "" {
		t.Fatalf(errMsg)
	}
}

func testTrim(str string, expected string, characterMask string) string {
	var msg string
	result := phpString.Trim(str, characterMask)
	want := regexp.MustCompile(expected)
	if !want.MatchString(result) {
		msg = fmt.Sprintf(`Trim("%v") = %q ,want match for %q`, str, result, want)
	}
	return msg
}
