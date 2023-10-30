package main

import "testing"

func Test(t *testing.T) {
	res := UnpackString("a4bc2d5e")
	if res != "aaaabccddddde" {
		t.Errorf("Expected aaaabccddddde, but got %s", res)
	}

	res = UnpackString("abcd")
	if res != "abcd" {
		t.Errorf("Expected abcd, but got %s", res)
	}

	res = UnpackString("45")
	if res != "[ERROR] - Некорректная строка" {
		t.Errorf("Expected [ERROR] - Некорректная строка, but got %s", res)
	}

	res = UnpackString("")
	if res != "" {
		t.Errorf("Expected , but got %s", res)
	}
}
