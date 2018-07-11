package main

import "fmt"

type rectangle struct{
	width,height int
}

func (r *rectangle) area() int{
	return r.height*r.width
}

func (r rectangle) perim() int{
	return 2*r.height*r.width
}

func Methods(){
	r:= rectangle{width:10,height:5}
	
	fmt.Println(r.area(),r.perim())

	rp := &r

	fmt.Println(rp.area(),rp.perim())
}

func main(){
	Methods()
}