package main

import (
	"fmt"
	"time"
)

type easyContext struct{
	ch chan struct{}
	done     chan struct{}
}
type easyface interface{
	easyWithCancel() (easyContext,func() )
	easyDone() <-chan struct{}
}
func (c *easyContext) easyDone(){
	d:=c.done
	return d
}

func (c *easyContext) easyWithCancel(c *easycontext){

	return c,func(){
		close(c.ch)
	}
}


func worker( ctx context, t *time.Ticker) {
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-ctx.easyDone():
				fmt.Println("recv stop signal")
				return
			case <-t.C:
				fmt.Println("is working")
			}
		}
	}()
	return
}

func main() {
	ticker := time.NewTicker(time.Second)
	p:=&easyContext{
		ch:make(chan struct{}),
	}
	var c easyface
	c=p
	ctx, cancel :=c.easyWithCancel(c)

	go worker(ctx, ticker)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(50 * time.Millisecond)
}