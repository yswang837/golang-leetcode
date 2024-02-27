package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type resultItem struct {
	Id    int
	Score float32
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

type race struct {
	total  int
	runner int
	wg     *sync.WaitGroup
	judge  chan struct{}
	result chan resultItem
}

func newRace(total, runner int) *race {
	return &race{
		total:  total,
		runner: runner,
		wg:     &sync.WaitGroup{},
		judge:  make(chan struct{}, 1),
		result: make(chan resultItem, runner),
	}
}

func main() {
	r := newRace(110, 9)
	r.judge <- struct{}{}
	time.Sleep(time.Second)
	<-r.judge
	for i := 0; i < r.runner; i++ {
		r.wg.Add(1)
		go r.run(i)
	}
	r.wg.Wait()
	close(r.judge)
	close(r.result)
	var raceResult []resultItem
	for ret := range r.result {
		raceResult = append(raceResult, ret)
	}
	sort.Sort(sortResult(raceResult))
	fmt.Printf("%d米跨栏比赛成绩如下:\n", r.total)
	for i, val := range raceResult {
		fmt.Printf("第%d名，编号是%d，耗时%.2fs\n", i+1, val.Id+1, val.Score)
	}
}

func (r *race) run(i int) {
	speed := 5 * (rand.Float32() + 2)
	r.result <- resultItem{i, float32(r.total) / speed}
	r.wg.Done()
}
