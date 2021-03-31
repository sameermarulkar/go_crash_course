package main

import (
	"fmt"
	"math"

	"github.com/sameermarulkar/go_crash_course/tree/master/03_packages/strutil"		// package imports not working - to be debugged
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}
