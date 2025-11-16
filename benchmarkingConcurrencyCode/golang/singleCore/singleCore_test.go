package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func init() {

	totalNumOfCores := runtime.NumCPU()
	numberOfCoresToUse := 1

	// force GO to use just the single core
	runtime.GOMAXPROCS(numberOfCoresToUse)

	// getter on GOMAXPROCS to get the current value; 0 is used as getter
	currentGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Printf("Total number of cores: %d \nGOMAXPROCS set to %d by using GOMAXPROCS(%d)\n", totalNumOfCores, currentGOMAXPROCS, numberOfCoresToUse)
}

func main() {
	fmt.Println("Starting the program")
}

func cpuWork() {
	total := 0
	for i := 0; i < 30_000_000; i++ {
		total += i
	}
}

func ioWork() {
	time.Sleep(500 * time.Millisecond)
}

// b.N is number of iterations; it is determined by the testing package;
// it is affected by the `benchtime` parameter used while running;

func BenchmarkSequentialCPUWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpuWork()
		cpuWork()
		cpuWork()
		cpuWork()
	}
}

func BenchmarkConcurrentSingleCoreCPUWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(4)

		go func() {
			defer wg.Done()
			cpuWork()
		}()
		go func() {
			defer wg.Done()
			cpuWork()
		}()
		go func() {
			defer wg.Done()
			cpuWork()
		}()
		go func() {
			defer wg.Done()
			cpuWork()
		}()

		wg.Wait()
	}
}

func BenchmarkSequentialIOWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioWork()
		ioWork()
		ioWork()
		ioWork()
	}
}

func BenchmarkConcurrentSingleCoreIOWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(4)

		go func() {
			defer wg.Done()
			ioWork()
		}()
		go func() {
			defer wg.Done()
			ioWork()
		}()
		go func() {
			defer wg.Done()
			ioWork()
		}()
		go func() {
			defer wg.Done()
			ioWork()
		}()

		wg.Wait()
	}

}
