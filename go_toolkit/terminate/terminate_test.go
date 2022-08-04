package terminate

import (
	"fmt"
	"testing"
	"time"
)

func TestTerminate(t *testing.T) {

	ch := make(chan struct{})
	fn := func() {
		fmt.Println("fnfnfnfn")
	}

	mc := new(map[int]int)
	fmt.Println(mc)

	go func() {
		select {
		case <-time.After(3 * time.Second):
			close(ch)
			fmt.Println("ch is closed...")
		}
	}()

	Terminate(ch, fn)
}
