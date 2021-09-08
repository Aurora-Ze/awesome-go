package main

import "fmt"

type test struct {
	val *Value
}

type Value struct {
	a int
}

func main() {
	cache := make(map[string]*test)
	cache["1"] = &test{
		val: &Value{
			a: 10,
		},
	}

	ele := cache["1"]
	ele.val.a = 10

	fmt.Println(ele.val.a)
}
