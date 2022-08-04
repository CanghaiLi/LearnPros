package terminate

import (
	"fmt"
	"time"
)

func Terminate(ch chan struct{}, fn func()) {
	for {
		select {
		// struct{}的通道在close时，会接受到值
		case <-ch:
			return
		case <-time.After(1 * time.Second):
			fmt.Println("task is running...")
			fn()
		}
	}
}
