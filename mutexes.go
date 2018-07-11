package main

import(
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Mutex()  {
	var stat = make(map[int]int)
	var metux = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for i := 0; i < 100; i++ {
		go func(){
			total:=0
			for{
				key:=rand.Intn(5)
				metux.Lock()
				total +=stat[key]
				metux.Unlock()
				atomic.AddUint64(&readOps,1)

				time.Sleep(time,Millsecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func ()  {
			for{
				key := rand.Intn(5)
				val :=rand.intn(100)
				metux.Lock()
				stat.Lock()
				metux.Unlock()
				atomic.AddUint64(&writeOps.1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
}