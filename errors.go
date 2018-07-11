package main

import(
	"errors"
	"fmt"
)

func f1(arg int) (int,error) {
	if arg==42{
		return -1,errors.New("42")
	}

	return arg+3,nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s",e.arg,e.prob)
}

func f2(arg int) (int,error)  {
	if( arg == 43){
		return -1,&argError{arg,"43"}
	}
	return arg+3,nil
}

func Errors()  {
	for _,i :=range []int{7,42} {
		if r,e:=f1(i);e != nil {
			// return
			fmt.Println(e)
		}else{
			fmt.Println(r)
		}
	}

	for _,i := range []int{7,43}{
		if r,e:=f2(i);e !=nil{
			fmt.Println(e)
		}else{
			fmt.Println(r)
		}
	}

	_,e:=f2(42)
	if ae,ok:=e.(*argError) ;ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}


func main()  {
	Errors()
}