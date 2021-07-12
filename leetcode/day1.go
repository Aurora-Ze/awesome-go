package leetcode

import (
	"fmt"
	"strings"
)

// Offer05：替换字符串中的空格为%20
func ReplaceSpace(s string) string {
	build := strings.Builder{}

	for _, v := range s {
		if v == ' ' {
			build.WriteString("%20")
		} else {
			build.WriteString(string(v))
		}
	}

	return build.String()
}

func Test() {
	str := "hello world"
	fmt.Println(str[0])
	fmt.Println(string(104))
	fmt.Println(str[5])
	fmt.Println(' ')
}
