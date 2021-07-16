package worker

import (
	"WorkPoolDemo/src/job"
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs chan job.Job) {
	defer wg.Done()
	for {
		select {
		case jo, ok := <-jobs:
			if !ok {
				return
			}
			jo.Execute(ctx)
		case <-ctx.Done():
			fmt.Printf("Worker canceled! Error : %v \n", ctx.Err())
			return
		}
	}

}

type WorkerPool struct {
	numsOfWorkers int
	//jobs是一个缓冲channel。每一个任务都会放入jobs中等待处理woker处理。
	jobs          chan job.Job
	Done          chan struct{}
}

func ConstructWorkerPool(num int) WorkerPool {
	return WorkerPool{
		numsOfWorkers: num,
		jobs:          make(chan job.Job, num),
		Done:          make(chan struct{}),
	}
}

func (pool WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < pool.numsOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, pool.jobs)
	}
	wg.Wait()
	close(pool.Done)

}

func (pool WorkerPool)AddJob(jo job.Job){
	pool.jobs <- jo
}
