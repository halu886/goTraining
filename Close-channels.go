package main

import "fmt"

func CloseChannels(){
	jobs:=make(chan int, 5)
	done:=make(chan bool, 1)

	go func(){
		for{
			j,more:=<-jobs
			if more{
				fmt.Println(j)
			} else{
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs<-i
	}

	close(jobs)

	<-done
}

func main()  {
	CloseChannels()
}