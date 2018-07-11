package main

import(
	"fmt"
	"time"
)

func RareLimiting()  {
	request:=make(chan int,5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	limit := time.Tick(500*time.Millisecond)

	for req := range request {
		<-limit
		fmt.Println(req,time.Now())
	}

	limitTicker := make(chan time.Time,5)

	for j := 0; j < 5; j++ {
		limitTicker <- time.Now()
	}

	go func()  {
		for ticker := range time.Tick(2000*time.Millisecond) {
			limitTicker <-ticker
		}
	}()

	burstRequest :=make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequest<-i
	}
	close(burstRequest)

	for req := range burstRequest {
		<-limitTicker
		fmt.Println(req,time.Now())
	}
}

func main()  {
	RareLimiting()
}
