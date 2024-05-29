package main

import (
	"fmt"
	"strings"
)

type MyString string

func (ms MyString) Capitalize() MyString {
	if len(ms) == 0 {
		return ms
	}

	return MyString(strings.ToUpper(string(ms[0]))) + MyString(ms[1:])
}
func main() {
	str := MyString("hello, world!")

	capitalized := str.Capitalize()
	fmt.Println(capitalized)
}