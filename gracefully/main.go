package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) Run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			selectJOb(strings.TrimSpace(os.Args[1]))
		}
	}
}

func (t *Task) Stop() {
	close(t.closed)
	t.wg.Wait()
}

func selectJOb(functionName string) {
	switch functionName {
	case "app1":
		app1()
	case "app2":
		app2()
	}
}

func app1() {
	for i := 0; i < 5; i++ {
		fmt.Print("@")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println()
}

func app2() {
	for i := 0; i < 5; i++ {
		fmt.Print("#")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println()
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 2),
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	task.wg.Add(1)
	go func() {
		defer task.wg.Done()
		task.Run()
	}()

	select {
	case sig := <-gracefulStop:
		fmt.Printf("Got %s signal. Aborting...", sig)
		task.Stop()
	}
}
