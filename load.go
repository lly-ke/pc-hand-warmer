package main

import (
	"math"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// LoadCPU cpu负载
func LoadCPU(model int) {
	numCPU := runtime.NumCPU()
	//runtime.GOMAXPROCS(numCPU)

	switch model {
	case 0:
		for i := 0; i < int(float64(numCPU)*0.6); i++ {
			wg.Add(1)
			go computePI()
		}
	case 1:
		for i := 0; i < int(float64(numCPU)*0.8); i++ {
			wg.Add(1)
			go computePI()
		}
	case 2:
		//协程好像跑步满
		for i := 0; i < int(float64(numCPU)*1); i++ {
			wg.Add(1)
			go computePI()
		}
	}
	wg.Wait()
}

// 模拟计算pi
func computePI() {
	var y = 1.0
	for {
		y = math.Sqrt(2 - math.Sqrt(4-y*y))
	}
}
