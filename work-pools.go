package main
import (
	"fmt"
	"time"
)

func worker1(id int,job <-chan int,result chan<- int)  {
	for j :=range job{
		fmt.Println(id,j)
		time.Sleep(time.Second)
		fmt.Println(id,j)
		result <- j*2
	}
}

func WorkerPool()  {
	jobs := make(chan int, 100)
	results := make(chan int,100)

	for w := 0; w < 3; w++ {
		go worker1(w,jobs,results)
	}

	for j := 0; j < 5; j++ {
		jobs<-j
	}
	close(jobs)

	for a := 0; a < 5; a++ {
		<-results
	}
}

func main()  {
	WorkerPool()
}