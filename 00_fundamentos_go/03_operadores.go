package main 

import ("fmt""reflect")

func main(){

	/* 
		Operadores 
		adição / concatenização = +
		subtração  = - 
		 multiplicação = * 
		divisão = / 
	*/
  num1 := 10.0
  num2 := 20.0

  result:= num1 / num2

  fmt.Println(result)
  fmt.Println(reflect.TypeOf(result)) // passara tipo do result


}