package main

import (
	"fmt"
	"time"
	"github.com/alcarruth/go-scratch/tinker"
)

func main() {
	fmt.Println("Howdy Al!")
	ctr := &tinker.Counter{}
	for {
		fmt.Printf("%d\n", ctr.Next())
		time.Sleep(1<<29)
	}
}
