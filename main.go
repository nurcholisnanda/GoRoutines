package main

import (
	"fmt"
	"goroutine/models"
	"goroutine/utils/csv"
	"strconv"
	"sync"
	"time"
)

var currentThread int

func main() {

	timeAllProcess := time.Now()
	//read file before Eod
	afterEodData, err := csv.ReadBeforeEodCsv("Before Eod.csv")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//use wait group to wait for each process to be done before continue to another process
	var wg sync.WaitGroup

	// processing process 1
	wg.Add(len(afterEodData))
	timeProcess1 := time.Now()
	currentThread = -1
	for i := 0; i < len(afterEodData); i++ {
		go process1(&wg, &afterEodData[i])
	}
	wg.Wait()
	fmt.Println("process 1 elapsed:", time.Since(timeProcess1))

	//processing process 2
	wg.Add(len(afterEodData))
	timeProcess2 := time.Now()
	currentThread = -1
	for i := 0; i < len(afterEodData); i++ {
		go process2(&wg, &afterEodData[i])
	}
	wg.Wait()
	fmt.Println("process 2 elapsed:", time.Since(timeProcess2))

	//processing process 3
	timeProcess3 := time.Now()
	totalData, maxThread := 100, 8
	for i := 0; i < totalData; i += maxThread {
		currentThread = -1
		limit := maxThread
		if i+maxThread > totalData {
			limit = totalData - i
		}
		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go process3(&wg, &afterEodData[j+i])
		}
		wg.Wait()
	}
	fmt.Println("process 3 elapsed:", time.Since(timeProcess3))

	//write updated data to after eod csv
	err = csv.WriteAfterEodCsv("After Eod.csv", afterEodData)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("all process elapsed:", time.Since(timeAllProcess))
}

//function for doing procces 1
func process1(wg *sync.WaitGroup, data *models.AfterEod) {
	defer wg.Done()
	currentThread++
	data.No1ThreadNo = strconv.Itoa(currentThread)
	data.AverageBalanced = (float32(data.Balanced) + float32(data.PreviousBalanced)) / float32(2)
}

//function for doing procces 2
func process2(wg *sync.WaitGroup, data *models.AfterEod) {
	defer wg.Done()
	currentThread++
	if data.Balanced >= 100 && data.Balanced <= 150 {
		data.FreeTransfer = 5
		data.No2AThreadNo = strconv.Itoa(currentThread)
	} else if data.Balanced > 150 {
		data.Balanced = data.Balanced + 25
		data.No2BThreadNo = strconv.Itoa(currentThread)
	}
}

//function for doing procces 3
func process3(wg *sync.WaitGroup, data *models.AfterEod) {
	defer wg.Done()
	currentThread++
	data.No3ThreadNo = strconv.Itoa(currentThread)
	data.Balanced = data.Balanced + 10
}
