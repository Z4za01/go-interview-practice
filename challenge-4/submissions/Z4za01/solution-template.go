package main

import (
	"fmt"
	"sync"
)
type result struct {
	start int
	order []int
}

type Queue []int

func (q *Queue) En(val int){
	*q = append(*q, val)
}
func (q *Queue) IsEmpty() bool{
	return len((*q)) == 0
}
func (q *Queue) De() (int, bool) {
    if len(*q) == 0 {
        return 0, false
    }
    v := (*q)[0]
    *q = (*q)[1:]
    return v, true
}

func bfsOne(graph map[int][]int, start int) []int {
    // TODO: 你來寫！
    // 提示：
    // 1) visited := map[int]bool{}
	visited := make(map[int]bool)
    // 2) queue := []int{start}
	var queue Queue
	queue.En(start)
	visited[start] = true
	order := []int{}
	for {
		v, ok := queue.De()
		if !ok {
			break
		}
		order = append(order, v)
		// fmt.Println(order)
		for _, element := range graph[v]{
			if !visited[element] {
				queue.En(element)
				visited[element] = true
			}
		}
	}

    return order
}
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	out := make(map[int][]int)
	if numWorkers < 1 {
		return out
	}
	if len(queries) == 0 {
		return out
	}
	jobs := make(chan int)
	results := make(chan result, len(queries)) 
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for start := range jobs {
				order := bfsOne(graph, start)
				results <- result{start: start, order: order}
			}
		}()
	}

	go func() {
		for _, q := range queries {
			jobs <- q
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		out[r.start] = r.order
	}

	return out
}

func main() {
	graph := map[int][]int{
        0: {1, 2},
        1: {2, 3},
        2: {3},
        3: {4},
        4: {},
    }
    queries := []int{0, 1, 2}
    numWorkers := 2
	bfsOne(graph, queries[2])
    results := ConcurrentBFSQueries(graph, queries, numWorkers)
	/*
       Possible output:
       results[0] = [0 1 2 3 4]
       results[1] = [1 2 3 4]
       results[2] = [2 3 4]
    */
	for i := 0;i < len(results);i++{
		fmt.Println(results[i])
	}
}
