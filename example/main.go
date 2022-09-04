// @author cold bin
// @date 2022/9/5

package main

import (
	"fmt"
	deep_copy "github.com/cold-bin/deep-copy"
)

type Struct struct {
	Int int
	Map map[int]string
}

func main() {
	s := Struct{
		Int: 1,
		Map: map[int]string{1: "1"},
	}
	newS := deep_copy.Copy(s)
	s.Map[1] = "2"
	fmt.Println(newS)
	fmt.Println(s)
}
