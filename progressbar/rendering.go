package progressbar

import (
	"fmt"
)

const bar = '█'

func (self *normalBar) String() string {
	return fmt.Sprintf("\r[%-20s]%3.f%%", string(self.bars()), self.Progress()*100)
}

func (self *normalBar) bars() []rune {
	bars := make([]rune, uint(self.Progress()*20))
	for i := range bars {
		bars[i] = bar
	}
	return bars
}
