package main

import "fmt"
import "time"

type Go struct {
	ChanCb    chan error
	pendingGo int
}

func New() *Go {
	a := new(Go)
	a.ChanCb = make(chan error, 1)
	return a
}

func (g *Go) Go(f func() error) {
	//f
	go func() {
		g.ChanCb <- f()
	}()

	for {
		select {
		case err := <-g.ChanCb:
			if err != nil {
				fmt.Println("error !!!")
			}
			return
		}
	}

}

func main() {
	a := New()
	a.Go(func() error {
		fmt.Println("start working...")
		time.Sleep(time.Second * 2)
		fmt.Println("end work!>")
		return nil
	})

}
