# Forwarder Pattern — Merge Worker Results

## Problem

You have multiple workers. Each worker produces results on its own receive-only channel and closes that channel when finished. Implement a `merge` function that combines results from all worker channels into a single channel.

## Worker signature

```go
func worker(id int) <-chan int
```

## Task

Implement:

```go
func merge(channels ...<-chan int) <-chan int
```

The returned channel should deliver all values produced by the provided worker channels and be closed once all input channels are closed.

## Example

Suppose the workers produce:

- `worker1` -> 1, 2, 3
- `worker2` -> 10, 20, 30
- `worker3` -> 100, 200, 300

Use `merge` like this:

```go
merged := merge(
    worker(1),
    worker(2),
    worker(3),
)

for v := range merged {
    fmt.Println(v)
}
```

The printed order may interleave values from different workers, for example:

```
1
10
100
2
20
200
3
30
300
```

## Notes

- The merge function should not leak goroutines.
- Preserve concurrency: read from input channels concurrently and forward values to the output channel.
- Close the output channel after all inputs are exhausted.

Good luck implementing the forwarder/merge pattern!