package main

import (
	"fmt"
	"progressbar/progressbar"
	"time"
)

func main() {
	bar := progressbar.New()
	for i := bar.Done(); i <= bar.Total(); i++ {
		bar.Set(i)
		fmt.Print(bar)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println()
}
