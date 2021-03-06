package progressbar

import (
	"fmt"
)

const BAR, NUM = '█', '#'

func (self *simple) String() string {
	return fmt.Sprintf("\r[%-20s]%3.f%%", string(self.bars()), self.Progress()*100)
}

func (self *simple) bars() []rune {
	bars := make([]rune, uint(self.Progress()*20))
	for i := range bars {
		bars[i] = self.bar
	}
	return bars
}
