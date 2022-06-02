package go2php

import (
	"fmt"
	"github.com/chenbo29/go2php/phpString"
	"regexp"
	"testing"
)

func TestSubstr(t *testing.T) {
	var str = "abcdef"
	if errMsg := testSubstr(str, "bcdef", 1, 0); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "bcd", 1, 3); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "cde", 2, -1); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "", 4, -4); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "abcd", 0, 4); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "abcde", 0, -1); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "de", -3, -1); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "ef", -2, 0); errMsg != "" {
		t.Fatalf(errMsg)
	}
	if errMsg := testSubstr(str, "d", -3, 1); errMsg != "" {
		t.Fatalf(errMsg)
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

func testSubstr(str string, expected string, start int64, length int64) string {
	msg := ""
	result, err := phpString.Substr(str, start, length)
	want := regexp.MustCompile(expected)
	if !want.MatchString(result) || err != nil {
		msg = fmt.Sprintf(`Substr("%v", "%b", "%b") = %q, %q ,want match for %q, nil`, str, start, length, result, err, want)
	}
	return msg
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
