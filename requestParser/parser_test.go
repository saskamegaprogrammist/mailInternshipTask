package requestParser

import (
	"testing"
)

// tests for working directory

const urlFirst = "aaa"
const errorFirst = "error getting url aaa request: Get \"aaa\": unsupported protocol scheme \"\""

func Test_URL_1(t *testing.T) {
	count, err := countURL(urlFirst)
	if count != 0 {
		t.Errorf("Test_URL_1 failed: count for %s should be 0", urlFirst)
	}
	if err == nil {
		t.Errorf("Test_URL_1 failed: error shouldn\\'t be nil")
	} else {
		if err.Error() != errorFirst {
			t.Errorf("Test_URL_1 failed: wrong error message")
		}
	}
}

const urlSecond = "https://golan"
const errorSecond = "error getting url https://golan request: Get \"https://golan\": dial tcp: lookup golan on 127.0.0.53:53: server misbehaving"

func Test_URL_2(t *testing.T) {
	count, err := countURL(urlSecond)
	if count != 0 {
		t.Errorf("Test_URL_2 failed: count for %s should be 0", urlSecond)
	}
	if err == nil {
		t.Errorf("Test_URL_2 failed: error shouldn\\'t be nil")
	} else {
		if err.Error() != errorSecond {
			t.Errorf("Test_URL_2 failed: wrong error message")
		}
	}
}

const fileFirst = "c.txt"
const errorFileFirst = "error reading file c.txt: stat c.txt: no such file or directory"

func Test_File_1(t *testing.T) {
	count, err := countFile(fileFirst)
	if count != 0 {
		t.Errorf("Test_File_1 failed: count for %s should be 0", fileFirst)
	}
	if err == nil {
		t.Errorf("Test_File_1 failed: error shouldn\\'t be nil")
	} else {
		if err.Error() != errorFileFirst {
			t.Errorf("Test_File_1 failed: wrong error message")
		}
	}
}

const fileSecond = "../testFiles/a.txt"
const countFileSecond = 8

func Test_File_2(t *testing.T) {
	count, err := countFile(fileSecond)
	if count != countFileSecond {
		t.Errorf("Test_File_2 failed: count for %s should be %d", fileSecond, countFileSecond)
	}
	if err != nil {
		t.Errorf("Test_File_2 failed: %s", err.Error())
	}
}

const fileThird = "../testFiles/b.txt"
const countFileThird = 64

func Test_File_3(t *testing.T) {
	count, err := countFile(fileThird)
	if count != countFileThird {
		t.Errorf("Test_File_3 failed: count for %s should be %d", fileThird, countFileThird)
	}
	if err != nil {
		t.Errorf("Test_File_3 failed: %s", err.Error())
	}
}
