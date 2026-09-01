package main

import (
	"fmt"
	"destd-lab/baseConv"
	"destd-lab/read"
)

func main() {
	var plaintext string = "10"
	fmt.Println(baseConv.Base2Conversion(plaintext))
	read.ReadInput()
}
