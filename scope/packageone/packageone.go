package packageone

import "fmt"

var PackageVar = "hello Package Var"

func PrintMe(p1, p2 string) {
	fmt.Println(p1, p2, PackageVar)
}
