package lab6

import (
	"PP_LABS/utils/mathutils"
	"PP_LABS/utils/stringutils"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func Init() {
	fmt.Println("Lab 6")
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
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
	numChl := make(chan int)
	strCh := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			numChl <- rand.Intn(10) + 1
			time.Sleep(time.Millisecond)
		}
		close(numChl)
	}()

	go func() {
		for {
			num := <-numChl
			if num%2 == 0 {
				strCh <- fmt.Sprintf("%d: even", num)
			} else {
				strCh <- fmt.Sprintf("%d: odd", num)
			}
		}
	}()

	for {
		select {
		case _, ok := <-numChl:
			if !ok {
				return
			}
		case parity, ok := <-strCh:
			if !ok {
				return
			}
			fmt.Println(parity)
		}
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

type calcParametr struct {
	id      int
	operand rune
	num1    float64
	num2    float64
	op      func(float64, float64) float64
}

func (c *calcParametr) Display() {
	fmt.Println("id: ", c.id, " | ", c.num1, string(c.operand), c.num2)
}

type calcResult struct {
	id  int
	res float64
}

func (c *calcResult) Display() {
	fmt.Println("id:", c.id, " | res: ", c.res)
}

func task5() {
	displayBorder()
	fmt.Println("Fifth task")
	numOp := 5
	wg := sync.WaitGroup{}
	wg.Add(numOp)
	res := make(chan *calcResult)
	par := make(chan *calcParametr)
	taskArr := []calcParametr{}
	for i := 0; i < numOp; i++ {
		taskArr = append(taskArr, *taskCreater(i))
	}
	for _, i := range taskArr {
		i.Display()
	}
	for i := 0; i < numOp; i++ {
		go func() {
			par <- &taskArr[i]
		}()
	}
	go func() {
		for i := 0; i < numOp; i++ {
			calculator(res, par, &wg)
		}
	}()
	for i := 0; i < numOp; i++ {
		(<-res).Display()
	}
	wg.Wait()
}

func taskCreater(id int) *calcParametr {
	randOp := rand.Intn(4) + 1
	randNum1 := rand.Intn(100) + 1
	randNum2 := rand.Intn(100) + 1
	switch randOp {
	case 1:
		return &calcParametr{id, '+', float64(randNum1), float64(randNum2), func(f1, f2 float64) float64 { return f1 + f2 }}
	case 2:
		return &calcParametr{id, '-', float64(randNum1), float64(randNum2), func(f1, f2 float64) float64 { return f1 - f2 }}
	case 3:
		return &calcParametr{id, '*', float64(randNum1), float64(randNum2), func(f1, f2 float64) float64 { return f1 * f2 }}
	case 4:
		return &calcParametr{id, '/', float64(randNum1), float64(randNum2), func(f1, f2 float64) float64 { return f1 / f2 }}
	}
	return &calcParametr{id, '+', 1.0, 1.0, func(f1, f2 float64) float64 { return f1 + f2 }}
}

func calculator(res chan<- *calcResult, task <-chan *calcParametr, wg *sync.WaitGroup) {
	defer wg.Done()
	c := <-task
	res <- &calcResult{c.id, c.op(c.num1, c.num2)}
}

func task6() {
	displayBorder()
	fmt.Println("Sixth task")
	file, err := os.Open("lab6_norm.txt")
	file2, err2 := os.OpenFile("lab6_rev.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	defer file.Close()
	defer file2.Close()
	var workers int = 0
	fmt.Println("Entered num of workers:")
	fmt.Fscan(os.Stdin, &workers)
	writeReverse(file, file2, workers)
}

func writeReverse(base *os.File, rev *os.File, workers int) {
	wg := sync.WaitGroup{}
	str := make(chan string, 100)
	res := make(chan string, 100)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go work(str, res, &wg, i)

	}
	scanner := bufio.NewScanner(base)
	for scanner.Scan() {
		str <- scanner.Text()
	}
	close(str)
	wg.Wait()
	close(res)
	for result := range res {
		rev.WriteString(result + "\n")
	}
}

func work(str <-chan string, rev chan<- string, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for s := range str {
		rev <- stringutils.Reversestring(s)
	}
}

func displayBorder() {
	fmt.Println("_______")
}
