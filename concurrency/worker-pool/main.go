package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int // job number (for logging)
	Number int // input number (e.g. to be squared)
}

type JobResult struct {
	Job    Job
	Result string
}

type worker[T, K any] struct {
	ID     int
	source chan T
	dest   chan K
}

func (w *worker[T, K]) run(runner func(T) K, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range w.source {
		result := runner(job)
		time.Sleep(1 * time.Second) // for visualizing
		w.dest <- result
	}
}

func newWorker[T Job, K JobResult](id int, src chan T, dest chan K) *worker[T, K] {
	return &worker[T, K]{ID: id, source: src, dest: dest}
}

func createNJobs(count int) []Job {
	jobs := make([]Job, 0, count)
	for i := 1; i <= count; i++ {
		jobs = append(jobs, Job{ID: i, Number: i})
	}

	return jobs
}

func runWorkers(
	count int,
	jobs chan Job,
	results chan JobResult,
	wg *sync.WaitGroup,
) {
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go newWorker(i, jobs, results).run(
			func(job Job) JobResult {
				result := JobResult{
					Job: job,
					Result: fmt.Sprintf(
						"Worker %d processed job %d: square(%d) = %d\n",
						i,
						job.ID,
						job.Number,
						job.Number*job.Number,
					),
				}

				return result
			},
			wg,
		)
	}
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan JobResult)

	// defer close(jobs)
	// defer close(results)

	for _, job := range createNJobs(10) {
		jobs <- job
	}
	close(jobs)

	var wg sync.WaitGroup

	runWorkers(2, jobs, results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r.Result)
	}
}
