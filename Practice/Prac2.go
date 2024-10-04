package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var db = [5]string{"Call1", "Call2", "Call3", "Call4", "Call5"}

var WG = sync.WaitGroup{}

func main() {
	fmt.Println("This is practice 2")
	for i := 0; i < len(db); i++ {
		WG.Add(1)
		go dbcall(i)
		WG.Wait()
		fmt.Printf("Total execution time: %v\n", time.Since(time.Now()))
	}

	var intslice = []int{1, 2, 3}
	fmt.Println(sum[int](intslice))

	var float32slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sum[float32](float32slice))

}

func sum[T int | float32](slice []T) T {
	var sum T
	for i := 0; i < len(slice); i++ {
		sum = sum + slice[i]
	}
	return sum
}

func dbcall(i int) {
	var delay = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("the result of the db is: %v", db[i])
	WG.Done()
}
