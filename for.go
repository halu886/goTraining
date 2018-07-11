package main

import "fmt"

func For()  {
	i:=1

	for i<=3{
		fmt.Println(i)
		i = i+1
	}

	for j:=8;j<10;j++{
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
	
	for n:=0;n<5;n++{
		if(n%2==0){
			continue
		}
		fmt.Println('n')
	}
}

func main () {
	For()
}