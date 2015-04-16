package main 

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input*2    
	fmt.Print(input,"*2=",output,"\n")


	C := (input - 32) * 5/9
	fmt.Println("Celsius= ",C)
}