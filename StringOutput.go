package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){

	fmt.Println("StringS")
	fmt.Println(strings.ToUpper("StringS"))
	
	reader :=bufio.NewReader(os.Stdin)
	fmt.Print("Enter text")
	line, _ := reader.ReadString('\n')
	
	fmt.Println(line)
	fmt.Println(strings.ToUpper(line))


}
