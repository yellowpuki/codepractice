package main

import (
	"bytes"
	"fmt"
)

func hasPrefix(str, prefix string) bool {

	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func hasSuffix(str, suffix string) bool {

	return len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix
}

func contain(str, substr string) bool {
	for i := 0; i < len(str); i++ {
		if hasPrefix(str[i:], substr) {
			return true
		}
	}

	return false
}

func countOfRunes(str string) int {
	n := 0
	for range str {
		n++
	}

	return n
}

func baseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}

	return path
}

func commaInt(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return commaInt(str[:n-3]) + "," + str[n-3:]
}

func commaFloat(str string) string {
	var (
		m string
		idx int
	)
	n:=len(str)
	for i:=n-1;i>=0;i--{
		if str[i] == '.' {
			idx = i
			m = str[i:]
			break
		}
	}
	return commaInt(str[:idx]) + m
}

func comma(str string) (out string) {
	var buf bytes.Buffer

	n:=len(str)
	for i := n-1; i>=0; i--{
		if i%3==0 && i != n-1 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(str[i]))
	}
	for _, s := range buf.String() {
		out = string(s) + out
	}
	return
}

func main() {
	str := "Hello World"
	prefix := "Hello"
	suffix := "World"
	substr := "o W"
  path := "/usr/local/.conf/conf.conf.a"

	fmt.Println(countOfRunes(str))
	fmt.Println(hasPrefix(str, prefix))
	fmt.Println(hasSuffix(str, suffix))
	fmt.Println(contain(str, substr))
	fmt.Println(contain(str, prefix))
	fmt.Println(contain(str, suffix))

	fmt.Println(baseName(path))

	fmt.Println(commaInt("1234567890"))
	fmt.Println(commaFloat("1234567890.1234567890"))
	fmt.Println(comma("1234567890"))
}
