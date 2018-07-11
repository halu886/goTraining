package main

import (
	"fmt"
)

func Maps()  {
	m:=make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m,m["k1"],len(m))

	delete(m,"k2")
	fmt.Println(m)

	_,prs := m["k1"]
	fmt.Println(prs)

	n:=map[string]int{
		"foo":1,
		"bar":2,
	}
	fmt.Println(n)
}

func main(){
	Maps()
}