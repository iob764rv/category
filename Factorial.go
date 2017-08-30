package main

import "fmt"

func factorial(a int) int{
	var result int;
	if a==0{
		return 1;
	}else{
		result=a*factorial(a-1)
		return result
	}
}

func main() {
	fmt.Println("Factorial is",factorial(4));
}

//factorials
