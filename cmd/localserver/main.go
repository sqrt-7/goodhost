package main

import (
	"fmt"

	"github.com/sqrt-7/goodhost/pkg/autogen"
)

func main() {
	stuff := &autogen.Item{}
	stuff.Id = "blah123"

	fmt.Println("ID:", stuff.Id)
}
