package main

import (
	"fmt"
	"time"
)

func main() {
	ctr := &Counter{}
	for {
		fmt.Printf("%d\n", ctr.next())
		time.Sleep(1<<29)
	}
}
