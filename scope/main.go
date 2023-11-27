package main

import (
	"main/packageone"
)

var myVar string = "mamamamama ma"

func main() {
	blockVar := "blockster"
	packageone.PrintMe(blockVar, myVar)
}
