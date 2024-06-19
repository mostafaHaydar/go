package main

import (
	"fmt"
	// "unicode"
)

func Concat(str1 string, str2 string) string {
	var newS string
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	runes := append(runes1, runes2...)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	newS = string(runes)
	return newS
}

func BasicJoin(elems []string) string {
	var newS string
	var runes []rune
	var tmpRunes []rune
	for i := 0; i < len(elems); i++ {
		tmpRunes = []rune(elems[i])
		runes = append(runes, tmpRunes...)
	}
	newS = string(runes)
	return newS
}
func Join(elems []string, sep string) string {
	var newS string
	var runes []rune
	var tmpRunes []rune
	sepRunes := []rune(sep)

	for i := 0; i < len(elems); i++ {
		tmpRunes = []rune(elems[i])
		runes = append(runes, tmpRunes...)
		runes = append(runes, sepRunes...)
	}
	newS = string(runes)
	return newS
}

func mssain() {
	elems := []string{"a", "b"}
	fmt.Printf("%s\n", Join(elems, "dd"))
}
