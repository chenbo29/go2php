package main

import (
	"fmt"
	"github.com/chenbo29/go2php/array"
)

func main() {
	str := []int{1, 2, 3}
	str = array.Push(str, 4)
	fmt.Println(str)
	strA := []string{"1", "2", "3"}
	strB := []string{"1", "2", "3"}
	strA = array.Merge(strA, strB)
	fmt.Println(strA)

	inArray := array.InArray(13, str)
	fmt.Println(inArray)
}
