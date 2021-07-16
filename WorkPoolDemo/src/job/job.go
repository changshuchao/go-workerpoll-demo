package job

import (
	"context"
	"fmt"
)

type Job struct {
	Name   int
	ExecFn ExecuteFunc
}

//type Result string

type ExecuteFunc func(ctx context.Context, args interface{}) (interface{}, error)


func (job Job) Execute(ctx context.Context) {
	_, err := job.ExecFn(ctx, job.Name)
	if err != nil {
		fmt.Printf("err : %v", err)
	}
}
