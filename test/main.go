package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type resultItem struct {
	Id    int
	Score time.Duration
}

type sortResult []resultItem

func (s sortResult) Len() int {
	return len(s)
}

func (s sortResult) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s sortResult) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	const runner = 10
	wg := &sync.WaitGroup{}
	result := make(chan resultItem, runner)
	judge := make(chan struct{}, 1)
	judge <- struct{}{}
	time.Sleep(time.Second)
	<-judge
	for i := 0; i < runner; i++ {
		wg.Add(1)
		go run(i, wg, result)
	}
	wg.Wait()
	close(judge)
	close(result)
	var raceResult []resultItem
	for r := range result {
		raceResult = append(raceResult, r)
	}
	sort.Sort(sortResult(raceResult))
	for i, val := range raceResult {
		fmt.Printf("第%d名，编号是%d,耗时%+v\n", i+1, val.Id, val.Score)
	}
}

func run(i int, wg *sync.WaitGroup, result chan resultItem) {
	start := time.Now()
	result <- resultItem{i, time.Since(start)}
	wg.Done()
}
