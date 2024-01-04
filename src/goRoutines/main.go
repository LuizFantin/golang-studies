package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	// goRoutines()
	goRoutinesMutex()

}

func goRoutines() {
	var wg sync.WaitGroup
	wg.Add(3)
	go actionA(&wg)
	go actionB(&wg)
	go actionC(&wg)

	wg.Wait()
}

func actionA(wg *sync.WaitGroup) {
	fmt.Println("A")
	wg.Done()
}

func actionB(wg *sync.WaitGroup) {
	fmt.Println("B")
	wg.Done()
}

func actionC(wg *sync.WaitGroup) {
	fmt.Println("C")
	wg.Done()
}

func goRoutinesMutex() {

	var m sync.Mutex
	number := 0
	for i := 0; i < 10000; i++ {
		go func() {
			m.Lock()
			number++
			m.Unlock()
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(number)

}
