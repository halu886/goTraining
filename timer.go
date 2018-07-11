package main
import (
	"fmt"
	"time"
)

func Timer()  {
	timer:= time.NewTimer(2*time.Second)

	<-timer.C
	fmt.Println("timer 1 expired")

	timer2:=time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("time 2 expired")
	}()

	stop2:=timer2.Stop()
	if(stop2){
		fmt.Println("stoped")
	}
}

func main()  {
	Timer()
}