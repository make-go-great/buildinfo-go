package main

import (
	"fmt"

	"github.com/make-go-great/buildinfo-go"
)

func main() {
	info, ok := buildinfo.Read()
	if !ok {
		return
	}

	fmt.Printf("%+v\n", info)
}
