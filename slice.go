package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func findDigit(filename string) []byte  {
	b,_ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte,len(b))
	c = append(c,b...)
	return digitRegexp.Find(b)
}

func Slices(){
	s:=make([]string,3)
	fmt.Println("emp",s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println(s,len(s))

	s = append(s,"d")
	s = append(s,"d","d")
	fmt.Println(s)

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println(c)
	
	l:=c[2:5]
	fmt.Println(l)

	l=c[:5]
	fmt.Println(l)
	
	l=c[2:]
	fmt.Println(l)

	twoD := make([][]int,3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int,innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)

	a:=[]string{"ad","dad"}
	b:=[]string{"asdf","dsaf"}
	a = append(a,b...)
	fmt.Println(a)
}

func main(){
	Slices()	
}