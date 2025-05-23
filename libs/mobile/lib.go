package mobile

import (
	"log"
	"sync"
)

func Run() {
	log.Println("Start service")
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
