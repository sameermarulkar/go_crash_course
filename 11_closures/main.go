// closures == annonymous functions

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	add := adder()				// here we are assigning the return response of adder i.e. func [annonymous function] to a variable add 
	
	fmt.Println("%T", add)

	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
