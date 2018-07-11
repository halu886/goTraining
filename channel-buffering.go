package main
import "fmt"

func ChannelsBuffering()  {
	message :=make(chan string,2)
	message <- "Buffered"	
	message <- "Channel"
	
	fmt.Println(<-message)
	fmt.Println(<-message)
}

func main()  {
	ChannelsBuffering()
}