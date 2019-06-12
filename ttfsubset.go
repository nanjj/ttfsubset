package main

import (
	"sort"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func GetAllRunes(fname string) (runes string, err error) {
	ttf, err := gofpdf.TtfParse(fname)
	if err != nil {
		return
	}
	chars := ttf.Chars
	max := len(chars)
	a := make([]string, max, max)
	i := 0
	for k := range chars {
		a[i] = string(rune(k))
		i++
	}
	sort.Strings(a)
	runes = strings.Join(a, "")
	return
}
