package core

import "sort"

func UniqueAppend(slice []string, value string) []string {
	for _, ele := range slice {
		if ele == value {
			return slice
		}
	}
	return append(slice, value)
}

func SortedKeys(m map[string][]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
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
