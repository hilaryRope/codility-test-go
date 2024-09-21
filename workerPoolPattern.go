package codility_test_go

import (
	"errors"
	"sync"
)

// WorkerPool errors, do not change!
var (
	ErrBadParams  = errors.New("bad params")
	ErrBadTask    = errors.New("bad task")
	ErrNotStarted = errors.New("not started")
)

// WorkerPool represents a pool of goroutines.
type WorkerPool struct {
	tasks   chan Task
	results chan error
	size    int
	started bool
	once    sync.Once
	wg      sync.WaitGroup
}

// Task to be computed by the WorkerPool.
type Task func() error

// NewWorkerPool creates a new pool with a given size.
func NewWorkerPool(size int) (*WorkerPool, error) {
	if size <= 0 {
		return nil, ErrBadParams
	}

	return &WorkerPool{
		tasks:   make(chan Task),
		results: make(chan error),
		size:    size,
	}, nil
}

// Results returns a channel of non-nil errors.
func (wp *WorkerPool) Results() <-chan error {
	return wp.results
}

// AddTask adds a task to the worker pool queue.
func (wp *WorkerPool) AddTask(task Task) error {
	if task == nil {
		return ErrBadTask
	}
	if !wp.started {
		return ErrNotStarted
	}
	wp.tasks <- task
	return nil
}

// Run will start workers (goroutines) for task computation.
func (wp *WorkerPool) Run() {
	wp.once.Do(func() {
		wp.started = true
		for i := 0; i < wp.size; i++ {
			wp.wg.Add(1)
			go wp.worker()
		}
		go func() {
			wp.wg.Wait()
			close(wp.results)
		}()
	})
}

// worker processes tasks and sends non-nil errors to the results channel.
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.tasks {
		if err := task(); err != nil {
			wp.results <- err
		}
	}
}