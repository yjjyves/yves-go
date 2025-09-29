package task

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

// WorkerPool 定义工作池
type WorkerPool struct {
	worker    int
	taskQueue chan func()
	wg        sync.WaitGroup
	closeOnce sync.Once
}

// NewWorkerPool 构造工作池
func NewWorkerPool(worker int, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		worker:    worker,
		taskQueue: make(chan func(), queueSize),
	}
	pool.Start()
	return pool
}

// Start 启动工作池
func (pool *WorkerPool) Start() {
	for i := 0; i < pool.worker; i++ {
		pool.wg.Add(1)
		go pool.Worker(i)
	}
}

// Worker 开始工作
func (pool *WorkerPool) Worker(id int) {
	defer pool.wg.Done()
	for task := range pool.taskQueue {
		logrus.Info("worker ", id, " start to process task")
		task()
	}
	logrus.Info("worker ", id, " finish process task")
}

func (pool *WorkerPool) Submit(task func()) {
	pool.taskQueue <- task
}

func (pool *WorkerPool) SubmitWithTimeout(ctx context.Context, task func()) error {
	select {
	case pool.taskQueue <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (pool *WorkerPool) Stop() {
	pool.closeOnce.Do(func() {
		close(pool.taskQueue)
		pool.wg.Wait()
	})

}
