package termcolor

import (
	"fmt"
	"strconv"
	"testing"
)

func TestReset(t *testing.T) {
	resetCode := Reset()
	if resetCode != start+strconv.Itoa(int(TermReset))+end {
		t.Fatalf("Reset code not ok")
	}
}

func TestGetCode(t *testing.T) {
	code := GetCode(Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s;%s;%s%s", start, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), end)
	if code != shouldbe {
		t.Fatal("Code Bold FgBlue BgBlue not ok")
	}
}

func TestFormat(t *testing.T) {
	code := Format("hello", Bold, FgBlue, BgBlue)
	shouldbe := fmt.Sprintf("%s%s%s;%s;%s%s%s%s", Reset(), start, strconv.Itoa(int(Bold)), strconv.Itoa(int(FgBlue)), strconv.Itoa(int(BgBlue)), end, "hello", Reset())
	if code != shouldbe {
		t.Fatal("Code Bold FgBlue BgBlue not ok")
	}
}
