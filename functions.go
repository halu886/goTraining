package main

import "fmt"

func plus(a int,b int) int  {
	return a+b
}

// func plus(a,b,c int) int { 
// 	return a+b+c
// }

func vals()(int,int){
	return 1,2
}

func sum (nums ...int) int{
	sum := 0
	for _,num := range nums{
		sum +=num
	}
	fmt.Println(sum)
	return sum
}

func Funtion(){
	fmt.Println(vals())
	_,c:=vals()
	fmt.Println(c)

	sum(1,3)
	sum(12,3,4)
	sum(([]int{12,3,5,6})...)
	// fmt.Println()
}

func main(){
	Funtion()
}

