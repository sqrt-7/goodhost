package main

import (
	"fmt"

	"github.com/sqrt-7/goodhost/pkg/autogen"
)

func main() {
	stuff := &autogen.MenuItem{}
	stuff.Id = "blah123"

	fmt.Println("ID:", stuff.Id)
}
