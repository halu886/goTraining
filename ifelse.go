package main

import "fmt"

func ElseIf()  {
	if(7%2==0){
		fmt.Println("a")
	}else{
		fmt.Println("b")
	}

	if 8%4==0{
		fmt.Println("c")
	}

	if num := 9;num<0{
		fmt.Println("a")
	} else if num<10{
		fmt.Println("b")
	} else{
		fmt.Println("c")
	}
}

func main(){
	ElseIf()
}