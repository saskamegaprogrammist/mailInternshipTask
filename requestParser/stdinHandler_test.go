package requestParser

import (
	"bytes"
	"strings"
	"testing"
)

// tests for working directory

const answerFirst = "Count for ../testFiles/a.txt: 8\nTotal: 8\n"
const inputFirst = "../testFiles/a.txt"

func Test_Simple_1(t *testing.T) {
	var errors []error
	var buf bytes.Buffer
	err := read(strings.NewReader(inputFirst), &buf, 5, &errors)
	if err != nil {
		t.Errorf("Test_Simple_1 failed: %s", err)
	}
	s := buf.String()
	if s != answerFirst {
		t.Errorf("Test_Simple_1 failed: wrong answer")
	}
}

const answerSecond = "Count for ../testFiles/a.txt: 8\nCount for ../testFiles/a.txt: 8\nCount for ../testFiles/b.txt: 64\nCount for ../testFiles/b.txt: 64\nCount for ../testFiles/b.txt: 64\nTotal: 208\n"
const inputSecond = "../testFiles/a.txt\n../testFiles/a.txt\n../testFiles/b.txt\n../testFiles/b.txt\n../testFiles/b.txt"

func Test_Simple_2(t *testing.T) {
	var errors []error
	var buf bytes.Buffer
	err := read(strings.NewReader(inputSecond), &buf, 5, &errors)
	if err != nil {
		t.Errorf("Test_Simple_2 failed: %s", err)
	}
	s := buf.String()
	if s != answerSecond {
		t.Errorf("Test_Simple_2 failed: wrong answer")
	}
}
