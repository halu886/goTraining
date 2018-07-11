package main

import "fmt"

func NoBlockingChannels()  {
	message := make(chan string)
	signals:=make(chan int)

	select{
	case msg:= <- message:
		fmt.Println(msg)
	default:
		fmt.Println("no received.")
	}

	msg:="hi"
	select{
	case message<-msg:
		fmt.Println("send message")
	default:
		fmt.Println("no message sent")
	}

	select{
	case msg:=<-message:
		fmt.Println(msg)
	case sig:=<-signals:
		fmt.Println(sig)
	default:
		fmt.Println("no actively")
	}
}

func main()  {
	NoBlockingChannels()
	
}