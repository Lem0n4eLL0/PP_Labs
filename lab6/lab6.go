package lab6

import (
	"PP_LABS/utils/mathutils"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Init() {
	fmt.Println("Lab 6")
	task1()
	task2()
	//task3()
	task4()
	task5()
}

func task1() {
	fmt.Println("First task")
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(n int64) {
		defer wg.Done()
		delay := rand.Intn(1000) + 10
		time.Sleep(time.Duration(delay))
		fmt.Println("factorial ", n, " = ", mathutils.Factorial(n), "\nEnd factorial func | delay is ", delay)
	}(6)

	go func(num ...float64) {
		defer wg.Done()
		delay := rand.Intn(1000) + 10
		time.Sleep(time.Duration(delay))
		var sum float64 = 0
		for _, n := range num {
			sum += n
		}
		fmt.Println("Sum ", num, " is ", sum, "\nEnd sum func | delay is ", delay)
	}(1, 2, 3)

	go func(lowerBound int, upperBound int) {
		defer wg.Done()
		delay := rand.Intn(1000) + 10
		time.Sleep(time.Duration(delay))
		fmt.Println("Rand num is ", rand.Intn(upperBound)+lowerBound, "\nEnd sum func | delay is ", delay)
	}(1, 100)
	wg.Wait()
}

func task2() {
	displayBorder()
	fmt.Println("Second task")
	wg := sync.WaitGroup{}
	wg.Add(2)
	intCh := make(chan int)
	go func(ch chan<- int) {
		defer close(ch)
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(10) + 1
		}
	}(intCh)

	go func(ch <-chan int) {
		defer wg.Done()
		for num := range ch {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}(intCh)
	wg.Wait()
}

func task3() {
	displayBorder()
	fmt.Println("Third task")
	intCh := make(chan int)
	stringCh := make(chan string)

	go func() {
		defer close(intCh)
		for i := 0; i < 1; i++ {
			intCh <- rand.Intn(10) + 1
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case res := <-intCh:
			numParity(res, stringCh)
		case res := <-stringCh:
			fmt.Println("Res: ", res)
		}
	}
}

func numParity(num int, ch chan string) {
	if num%2 == 0 {
		ch <- "even"
	} else {
		ch <- "add"
	}
}

var counter int = 0

func task4() {
	displayBorder()
	fmt.Println("Fourth task")
	var mutex sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go chageCounter(i, &mutex, &wg)
	}
	wg.Wait()
}

func chageCounter(n int, m *sync.Mutex, wg *sync.WaitGroup) {
	m.Lock()
	defer wg.Done()
	counter = 0
	for i := 0; i < 5; i++ {
		counter++
		fmt.Println("gn ", n, " counter ", counter)
	}
	m.Unlock()
}

func task5() {
	numOp := 5
	wg := sync.WaitGroup{}
	wg.Add(numOp)
	for i := 0; i < numOp; i++ {
		go taskCalc()
	}
	wg.Wait()
}
func taskCalc() {
	randOp := rand.Intn(4) + 1
	//randNum1 := rand.Intn(100) + 1
	//randNum2 := rand.Intn(100) + 1
	switch randOp {
	case 1:
		//calculator()
	case 2:
	case 3:
	case 4:
	}

}

func calculator(res chan<- float64, task <-chan func(float64, float64) float64, wg *sync.WaitGroup) {
	defer wg.Done()
	res <- (<-task)(1, 2)
}

func displayBorder() {
	fmt.Println("_______")
}
