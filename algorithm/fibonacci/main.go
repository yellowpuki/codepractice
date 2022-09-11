package main

import (
	"fmt"

	"github.com/yellowpuki/codepractice/advanced"
	"github.com/yellowpuki/codepractice/base"
)

func main() {
	fmt.Println(base.Fibo(5))
	fmt.Println(advanced.CachedFibo(20))
}
