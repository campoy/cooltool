package server

import (
	"fmt"
	"time"
)

func Run() {
	for i := 0; true; i++ {
		timestr := "seconds"
		if i == 1 {
			timestr = "second"
		}
		fmt.Println("running for", i, timestr, "and still didn't crash")
		time.Sleep(1 * time.Second)
	}
}
