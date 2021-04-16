/*
负责人员：InBug Team
创建时间：2021/4/16
程序用途：并发执行单元
*/
package coroutine

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Task func(*sync.WaitGroup, chan bool, chan int, chan interface{}, ...interface{})
type Iter func(*sync.WaitGroup, chan bool, chan int, chan interface{}, Task, ...interface{})
type Done func(interface{}) error
type Ing func(float64)
type After func(int)

func MultiTask(
	totalTask, workerNumber int,
	iter Iter,
	task Task,
	done Done,
	ing Ing,
	after After,
	msgStart, msgIng, msgEnd string,
	data ...interface{},
) {
	if totalTask == 0 {
		return
	}
	fmt.Println("****************<-START->****************")
	fmt.Println(msgStart)
	start := time.Now()
	if totalTask <= workerNumber {
		workerNumber = totalTask
	}
	var wg sync.WaitGroup
	worker := make(chan bool, workerNumber)
	counts := make(chan int)
	result := make(chan interface{}, workerNumber)
	ingTask, scaleTask, doneTask := 0, 0, 0
	size := int(math.Floor(float64(totalTask) / float64(10)))

	go func() {
		iter(&wg, worker, counts, result, task, data...)
		wg.Wait()
		close(worker)
	}()

	for {
		select {
		case number := <-counts:
			ingTask += number
			scaleTask = int((float64(ingTask) / float64(totalTask)) * 100)
			switch ingTask {
			case 1, size * 1, size * 2, size * 3, size * 4, size * 5, size * 6, size * 7, size * 8, size * 9, totalTask:
				ing(float64(scaleTask))
			}
			fmt.Println(fmt.Sprintf(`%s-进度%d%%，%d/%d`, msgIng, scaleTask, ingTask, totalTask))
			if ingTask == totalTask {
				after(doneTask)
				fmt.Println(fmt.Sprintf(`%s，执行总耗时：%f秒`, msgEnd, time.Since(start).Seconds()))
				fmt.Println("****************<-END->****************")
				return
			}
		case item := <-result:
			if err := done(item); err == nil {
				doneTask++
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
