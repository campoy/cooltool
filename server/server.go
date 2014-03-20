package server

import (
	"fmt"
	"time"
)

func Run() {
	for i := 0; true; i++ {
		fmt.Println("running for", i, "seconds and still didn't crash")
		time.Sleep(1 * time.Second)
	}
}
