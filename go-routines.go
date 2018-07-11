package main
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from,i)
	}
}

func GoRoutines()  {
	f("direct")

	go f("goRountine")

	go func (msg string)  {
		fmt.Println(msg)
	}("going")

	time.Sleep(2*time.Second)

	fmt.Println("Done")
}

func main()  {
	GoRoutines()
}