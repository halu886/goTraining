package main

import "fmt"

func Range()  {
	nums := []int{1,2,3}
	sum := 0
	for _,num :=range nums{
		sum +=num
	}
	fmt.Println(sum)


	for index,num := range nums{
		if(index==2){
			fmt.Println(num)
		}
	}

	kvs:= map[string]string{
		"a":"apple",
		"b":"banana",
	}
	for k,v:=range kvs{
		fmt.Printf("%s=>%s",k,v)
	}
	for i,c :=range "go"{
		fmt.Println(i,c)
	}
}

func main(){
	Range()
}