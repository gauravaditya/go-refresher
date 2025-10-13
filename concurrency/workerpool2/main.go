package main

import (
	"context"
	"fmt"
	"sync"
)

/*
Task: Implement a concurrent WorkerPool in Go that:
Has a constructor NewWorkerPool(concurrency int) that returns a *WorkerPool.
Exposes Submit(job func() error) error to enqueue a job. If the pool is shut down, Submit should return an error.
Exposes Shutdown(ctx context.Context) error which waits for running jobs to finish and returns when complete or when context is done.
The WorkerPool should accept up to a configurable buffer of pending jobs (choose a sensible default like 1024), and Submit should block if the buffer is full (backpressure).
If a job returns an error, the WorkerPool should capture the first error and return it from Shutdown (but still finish in-flight jobs).

Requirements:
- No goroutine leaks when Shutdown completes.
- Use context properly for shutdown waiting.
*/

func main() {
	pool := NewWorkerPool(3)

	var err error
	for i := range 100 {
		err = pool.Submit(func() error {
			fmt.Println(i)
			return nil
		})
	}
	fmt.Println("error received on pool.Submit:", err)

	err = pool.Shutdown(context.Background())
	fmt.Println("error received on pool.Shutdown:", err)
	err = pool.Submit(func() error {
		fmt.Println("Pushing value after shutdown...")
		return nil
	})
	fmt.Println("error received on pool.Submit after pool.Shutdown:", err)
}

func NewWorkerPool(concurrency int) *workerPool {
	pool := &workerPool{
		jobs:   make(chan func() error, 1024),
		closed: make(chan struct{}),
		wg:     sync.WaitGroup{},
	}

	for range concurrency {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

type workerPool struct {
	jobs    chan func() error
	closed  chan struct{}
	wg      sync.WaitGroup
	once    sync.Once
	errOnce sync.Once
	err     error
}

func (w *workerPool) worker() {
	defer w.wg.Done()
	for job := range w.jobs {
		if err := job(); err != nil {
			w.errOnce.Do(func() { // capture the first error only
				w.err = err
			})
		}
	}
}

func (w *workerPool) Submit(job func() error) error {
	select {
	case <-w.closed:
		return fmt.Errorf("cannot push, pool is closed..")
	default:
		w.jobs <- job
		return nil
	}
}

func (w *workerPool) Shutdown(ctx context.Context) error {
	select {
	case <-w.closed:
		return fmt.Errorf("pool already closed...")
	default:
	}

	w.once.Do(func() { // close channels only once
		close(w.closed)
		close(w.jobs)
	})

	done := make(chan struct{})
	go func() {
		w.wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("Shutdown called: %w", ctx.Err())
	case <-done:
		return nil
	}
}
