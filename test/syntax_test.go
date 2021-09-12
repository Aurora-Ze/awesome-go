package test

import (
	"log"
	"strings"
	"testing"
)

func Test_Map_NotFound(t *testing.T) {
	testMap := make(map[string]int)
	testMap["0"] = 1
	v, ok := testMap["1"]
	log.Printf("v is %v\n", v)
	log.Printf("ok is %v\n", ok)

	v1, ok1 := testMap["0"]
	log.Printf("v1 is %v\n", v1)
	log.Printf("ok1 is %v\n", ok1)
}

func Test_String_Split(t *testing.T) {
	str := "///"
	res := strings.Split(str, "/")
	log.Printf("res is %v\n", res)
	log.Printf("res length is %v\n", len(res))
	log.Printf("res[0] eq \"\": %v\n", res[0] == "")

}
