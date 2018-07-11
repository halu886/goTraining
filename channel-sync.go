package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Println("asdf")
	time.Sleep(time.Second)
	done<-true
	
}

func  ChannelSync()  {
	done:=make(chan bool,1)
	go worker(done)

	<-done
}

func main()  {
	ChannelSync()
}