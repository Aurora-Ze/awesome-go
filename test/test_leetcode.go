package test

import (
	"bytes"
	"fmt"
)

// Offer 05
func ReplaceSpace(s string) string {
	buffer := bytes.Buffer{}

	for _, v := range s {
		if v == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteString(string(v))
		}
	}

	return buffer.String()
}

func Test() {
	str := "hello world"
	fmt.Println(str[0])
	fmt.Println(string(104))
	fmt.Println(str[5])
	fmt.Println(' ')
}
