package main

import (
	"fmt"
	"time"
)

func Ticker(){
	ticker :=time.NewTicker(500*time.Millisecond)
	go func(){
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("stop")
}

func main()  {
	Ticker()
}