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
	var temp int
	if a<b{
		temp=b
		b=a
		a=temp
	}


}
