package progressbar_test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	. "progressbar/progressbar"
	"testing"
)

func TestInitialValues(t *testing.T) {
	bar := New(BAR)

	assert.Equal(t, uint(100), bar.Total())
	assert.Equal(t, uint(0), bar.Done())
	assert.Equal(t, float32(0), bar.Progress())
}

func TestSetDone(t *testing.T) {
	bar := New(BAR)

	bar.Set(20)

	assert.Equal(t, uint(100), bar.Total())
	assert.Equal(t, uint(20), bar.Done())
	assert.Equal(t, float32(0.2), bar.Progress())
}

func TestSetBackwardDoneWillBeSkipped(t *testing.T) {
	bar := New(BAR)

	bar.Set(20)

	bar.Set(10)
	assert.Equal(t, uint(100), bar.Total())
	assert.Equal(t, uint(20), bar.Done())
	assert.Equal(t, float32(0.2), bar.Progress())
}

func TestSetDoneOutOfTotalRangeWillMarkAsCompleted(t *testing.T) {
	bar := New(BAR)

	bar.Set(200)

	assert.Equal(t, uint(100), bar.Total())
	assert.Equal(t, uint(100), bar.Done())
	assert.Equal(t, float32(1), bar.Progress())
}

func TestPercentageOfProgressPresentation(t *testing.T) {
	bar := New(BAR)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `  0%$`)

	bar.Set(1)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `  1%$`)

	bar.Set(10)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `(?:^|[^ ]) 10%$`)
}

func TestBarPresentation(t *testing.T) {
	bar := New(BAR)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `\[                    \]`)

	bar.Set(1)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `\[                    \]`)

	bar.Set(10)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `\[██                  \]`)
}

func TestCustomBarPresentation(t *testing.T) {
	bar := New(NUM)

	bar.Set(10)
	assert.MatchRegex(t, bar.(fmt.Stringer).String(), `\[##                  \]`)
}
