package gosnippet

import (
	"fmt"
	"sync"
)

func main() {
	var (
		a = []interface{}{1, 2, 3, 4, 5, 6}
		b = []interface{}{"a", "b", "c", "d", "e"}

		wg sync.WaitGroup

		cha = make(chan struct{}, 1)
		chb = make(chan struct{}, 1)
	)

	wg.Add(2)

	go Run(&wg, a, cha, chb)
	go Run(&wg, b, chb, cha)

	chb <- struct{}{}
	wg.Wait()
}

func Run(group *sync.WaitGroup, data []interface{}, selfChan chan struct{}, peerChan chan struct{}) {
	for _, d := range data {
		if _, ok := <-peerChan; !ok {
			fmt.Println(d)
			continue
		}
		fmt.Println(d)
		selfChan <- struct{}{}
	}
	close(selfChan)
	group.Done()
}
