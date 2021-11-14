package progressbar_test

import (
	"fmt"
	"progressbar/progressbar"
)

func ExampleProgressBarStart() {
	bar := progressbar.New(progressbar.BAR)
	fmt.Print(bar)
	// output:
	// [                    ]  0%
}

func ExampleProgressBarDone() {
	bar := progressbar.New(progressbar.BAR)
	bar.Set(100)
	fmt.Print(bar)
	// output:
	// [████████████████████]100%
}
