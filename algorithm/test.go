package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	game := NewGame()
	// 10个运动员参加比赛
	for i := 1; i <= 10; i++ {
		runner := NewRunner(i)
		game.SetRunner(runner)
	}
	// 起跑线前准备
	game.Prepare()
	// 发令枪响
	game.Start()
	// 等待运动员跑完
	game.wg.Wait()
}

type Runner struct {
	id int
}

func (r *Runner) Speed() float64 {
	return 5 + rand.Float64()*10
}

func NewRunner(id int) *Runner {
	return &Runner{id: id}
}

type Game struct {
	length  int // 赛道
	runners []*Runner

	wg    sync.WaitGroup // 控制运动员是否全部跑完
	start chan struct{}  // 发令枪
	end   chan int       // 终点
}

func NewGame() *Game {
	return &Game{
		length:  100,
		runners: make([]*Runner, 0),
		start:   make(chan struct{}),
		end:     make(chan int),
	}
}

func (g *Game) SetLength(length int) {
	g.length = length
}

func (g *Game) SetRunner(runner *Runner) {
	g.runners = append(g.runners, runner)
}

func (g *Game) Prepare() {
	for _, runner := range g.runners {
		g.wg.Add(1)
		go func(r *Runner) {
			defer g.wg.Done()
			for {
				select {
				case <-g.start:
					time.Sleep(time.Duration(float64(g.length)/r.Speed()) * time.Second)
					g.end <- r.id
					return
				}
			}

		}(runner)
	}
	// 统计成绩
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		// 发令枪开始后才统计成绩
		for {
			select {
			case <-g.start: // 比赛开始，开始计时
				now := time.Now()
				top := 1
				for id := range g.end {
					spend := time.Now().Sub(now).Seconds()
					fmt.Printf("第%d名: \t运动员编号：%d\t花费时间：%f\n", top, id, spend)
					top++
					if top == len(g.runners)+1 {
						return
					}
				}
			}
		}
	}()
}

func (g *Game) Start() {
	close(g.start)
}
