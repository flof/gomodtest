package main

import (
	"fmt"

	"github.com/flof/gomodtest/sdk/v2"
)

func main() {
	s := sdk.Create(true)
	fmt.Println(s)
}
