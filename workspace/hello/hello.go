package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello, Go!"))
	fmt.Println(stringutil.ToUpper("Hello, Go!"))
}