package main

import "fmt"

func Channels()  {
	message := make(chan string)

	go func()  {
		message <- "ping"
	}()

	msg := <- message
	fmt.Println(msg)
}

func main()  {
	Channels()
}