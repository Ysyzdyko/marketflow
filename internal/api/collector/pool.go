package collector

import (
	"errors"
	"sync"
)

type WorkerPool struct {
	isClosed bool
	mu       sync.Mutex
	taskCH   chan func()
	doneCH   chan struct{}
}

func New(workerNums int) (*WorkerPool, error) {
	if workerNums <= 0 {
		return nil, errors.New("incorect WorkersNum")
	}

	wp := &WorkerPool{
		isClosed: false,
		mu:       sync.Mutex{},
		taskCH:   make(chan func(), 5),
		doneCH:   make(chan struct{}),
	}

	go wp.Proccess(workerNums)
	return wp, nil
}

func (wp *WorkerPool) Proccess(workersNums int) {
	wg := sync.WaitGroup{}
	wg.Add(workersNums)

	for i := 0; i < workersNums; i++ {
		go func() {
			defer wg.Done()
			for task := range wp.taskCH {
				task()
			}
		}()
	}

	wg.Wait()
	close(wp.doneCH)
}

func (wp *WorkerPool) AddTask(task func()) error {
	if task == nil {
		return errors.New("task is incorect")
	}

	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.isClosed {
		return errors.New("worler pool is closed ")
	}

	select {
	case wp.taskCH <- task:
		return nil
	default:
		return errors.New("workerPull is full")
	}
}

func (wp *WorkerPool) Close() error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.isClosed {
		return errors.New("worker pool is closed ")
	}

	wp.isClosed = true
	close(wp.taskCH)
	<-wp.doneCH
	return nil
}
