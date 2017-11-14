//lcd euclidean algorithm
package main

import (
	"fmt"
)

func main() {
	var num = 3
	fmt.Println(num)
	euclideanDenominator(0,1)

}

func euclideanDenominator(a int,b int){

	fmt.Println("function")
//	var i=a-b
	for a-b!=0{ //while{

	if a-b==0 || a==b{
	  fmt.Println("Denominator is ", b) }

	if a<b{
		b, a = a, b //swap elements }

  }
  }
}
