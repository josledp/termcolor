package termcolor

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetCode(t *testing.T) {
	code := GetCode(Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s;%s;%s%s", normalStart, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), normalEnd)
	if code != shouldbe {
		t.Fatal("Code Bold FgBlue BgBlue not ok")
	}
}

func TestFormat(t *testing.T) {
	code := Format("hello", Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s%s;%s;%s%s%s%s", GetCode(TermReset), normalStart, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), normalEnd, "hello", GetCode(TermReset))
	if code != shouldbe {
		t.Fatal("Code Bold FgBlue BgBlue not ok")
	}
}

func TestGetEscapedCode(t *testing.T) {
	code := GetEscapedCode(Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s;%s;%s%s", escapedStart, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), escapedEnd)
	if code != shouldbe {
		t.Fatal("Escoped Code Bold FgBlue BgBlue not ok")
	}
}

func TestEscapedFormat(t *testing.T) {
	code := EscapedFormat("hello", Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s%s;%s;%s%s%s%s", GetEscapedCode(TermReset), escapedStart, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), escapedEnd, "hello", GetEscapedCode(TermReset))
	if code != shouldbe {
		t.Fatal("Escoped Code Bold FgBlue BgBlue not ok")
	}
}
