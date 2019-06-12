package main

import "testing"

func TestGetAllRunes(t *testing.T) {
	runes, err := GetAllRunes("testdata/zhong.ttf")
	if err != nil {
		t.Fatal(err)
	}
	if runes != "中好您文！，" {
		t.Fatal(runes)
	}
}
