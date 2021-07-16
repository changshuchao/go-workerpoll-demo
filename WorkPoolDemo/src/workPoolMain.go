package main

import (
	"WorkPoolDemo/src/job"
	"WorkPoolDemo/src/worker"
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//go协程 大概耗时750ms
	//test()
	//workerpool形式 大概耗时650ms
	testPool()
}

func testPool(){
	begin := time.Now()
	//numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(100)
	//fmt.Println("cores : ",numCores)

	pool := worker.ConstructWorkerPool(100)
	ctx,cancle := context.WithCancel(context.TODO())
	defer cancle()
	go pool.Run(ctx)

	excuFn := func(ctx context.Context, args interface{}) (interface{}, error){
		sum(args.(int))
		return nil,nil
	}
	for i := 1;i<100000;i++{
		pool.AddJob(job.Job{
			Name: i,
			ExecFn: excuFn,
		})
	}
	end := time.Now()
	fmt.Println("cost time: %v",end.Sub(begin).Milliseconds())
}

func test() {
	begin := time.Now()
	var wg sync.WaitGroup
	for i := 1;i<100000;i++{
		wg.Add(1)
		go sum1(i,&wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("cost time: %v",end.Sub(begin).Milliseconds())
}

func sum1(n int,wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	//fmt.Println("%s 的累加是 %s", n, result)
}

func sum(n int) {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	//fmt.Println("%s 的累加是 %s", n, result)
}
