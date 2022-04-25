package core

import (
	"sort"
	"strings"
)

func UniqueAppend(slice []string, value string) []string {
	for _, ele := range slice {
		if ele == value {
			return slice
		}
	}
	return append(slice, value)
}

func SortedKeys(m map[string][]string, isFive bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	index := strings.Index(keys[0], "-")
	if index > 0 && isFive {
		var key = keys[0]
		var number = key[:(index)]
		if number == "10" {
			keys = append(keys[1:], keys[0])
		}
	}
	return keys
}

var NetflixCategory = newCategoryRegister()

func newCategoryRegister() *categoryEnum {
	return &categoryEnum{
		IlkUc:      1,
		IkinciUc:   2,
		UcuncuUc:   3,
		DorduncuUc: 4,
		BesinciUc:  5,
		IlkBes:     1,
		IkinciBes:  2,
		UcuncuBes:  3,
	}
}

type categoryEnum struct {
	IlkUc      int
	IkinciUc   int
	UcuncuUc   int
	DorduncuUc int
	BesinciUc  int
	IlkBes     int
	IkinciBes  int
	UcuncuBes  int
}
