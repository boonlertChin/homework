package main

import "fmt"
import "homework/boonlert/homework2"

func main() {
    
	testInput := []int{1, 2, 3, 4, 5} 
	k:= 12
	fmt.Println("n = ",homework2.IsTailSum(testInput,k))
}