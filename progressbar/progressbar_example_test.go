package progressbar_test

import (
	"fmt"
	"progressbar/progressbar"
)

func ExampleProgressBarStart() {
	bar := progressbar.New()
	fmt.Print(bar)
	// output:
	// [                    ]  0%
}

func ExampleProgressBarDone() {
	bar := progressbar.New()
	bar.Set(100)
	fmt.Print(bar)
	// output:
	// [████████████████████]100%
}
