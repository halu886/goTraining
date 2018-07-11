package main

import (
	"fmt"
	"time"
)

func Switch()  {
	i:=2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("weekenday")
	default:
		fmt.Println("weekDay")
	}

	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("it`s noon")
	default:
		fmt.Println("afterNoon")
	}

	whatAmI := func(i interface{}){
		switch t:=i.(type){
		case bool :
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("%T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(234)
	whatAmI("hey")
}

func main(){
	Switch()
}