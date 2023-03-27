package main

import (
	"fmt"
	"gvb_server/utils/random"
)

func main() {
	s := random.RandString(16)
	fmt.Println(s)
}
