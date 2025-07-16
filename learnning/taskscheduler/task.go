package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Task func()

type TaskContext struct {
	Name string
	Task Task
}

type TaskScheduler struct {
	Tcs []TaskContext
}

func (t *TaskScheduler) AddTask(name string, f func()) {
	tc := TaskContext{
		Name: name,
		Task: f,
	}

	t.Tcs = append(t.Tcs, tc)
}

var wg sync.WaitGroup

func (t TaskScheduler) Run() {
	tcSlice := t.Tcs
	for _, v := range tcSlice {
		wg.Add(1)

		go func(tc TaskContext) {
			startTime := time.Now().UnixMilli()
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					fmt.Errorf("任务执行失败，失败原因:%v,任务名:%v", r, v.Name)
				}
			}()
			v.Task()
			endTime := time.Now().UnixMilli()
			fmt.Printf("任务: %s 执行了 %v\n", v.Name, endTime-startTime)
		}(v)
	}
	wg.Wait()
}

func main() {
	taskScheduler := TaskScheduler{}
	for i := 1; i < 10; i++ {
		taskScheduler.AddTask(strconv.Itoa(i), func() {
			time.Sleep(time.Duration(1) * 1000)
		})
	}

	taskScheduler.Run()

}
