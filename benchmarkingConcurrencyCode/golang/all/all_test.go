package main

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func displayCoreInfo(numberOfCoresToUse int) {
	runtime.GOMAXPROCS(numberOfCoresToUse)
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

func BenchmarkSequentialSingleCoreCPUWork(b *testing.B) {
	displayCoreInfo(1)
	for i := 0; i < b.N; i++ {
		cpuWork()
		cpuWork()
		cpuWork()
		cpuWork()
	}
}

func BenchmarkConcurrentSingleCoreCPUWork(b *testing.B) {
	displayCoreInfo(1)
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

func BenchmarkSequentialSingleCoreIOWork(b *testing.B) {
	displayCoreInfo(1)
	for i := 0; i < b.N; i++ {
		ioWork()
		ioWork()
		ioWork()
		ioWork()
	}
}

func BenchmarkConcurrentSingleCoreIOWork(b *testing.B) {
	displayCoreInfo(1)
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

func BenchmarkSequentialMultiCoreCPUWork(b *testing.B) {
	displayCoreInfo(8)
	for i := 0; i < b.N; i++ {
		cpuWork()
		cpuWork()
		cpuWork()
		cpuWork()
	}
}

func BenchmarkConcurrentMultiCoreCPUWork(b *testing.B) {
	displayCoreInfo(8)
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

func BenchmarkSequentialMultiCoreIOWork(b *testing.B) {
	displayCoreInfo(8)
	for i := 0; i < b.N; i++ {
		ioWork()
		ioWork()
		ioWork()
		ioWork()
	}
}

func BenchmarkConcurrentMultiCoreIOWork(b *testing.B) {
	displayCoreInfo(8)
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
