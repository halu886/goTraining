package main
import (
	"fmt"
	"sync/atomic"
	"time"
)

func AtomicCounter()  {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func(){
			for{
				atomic.AddUint64(&ops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(5*time.Second)
	opsFinal:=atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)
}

func main()  {
	AtomicCounter()
}
