package progressbar

import . "sync/atomic"

type ProgressBar interface {
	Total() uint
	Done() uint
	Progress() float32
	Set(done uint)
}

type simple struct {
	done uint32
}

func (self *simple) Total() uint {
	return 100
}

func (self *simple) Done() uint {
	return uint(LoadUint32(&self.done))
}

func (self *simple) Progress() float32 {
	return float32(self.Done()) / float32(self.Total())
}

func (self *simple) Set(done uint) {
	for {
		current := self.Done()
		if current >= done {
			return
		}
		if CompareAndSwapUint32(&self.done, uint32(current), uint32(min(done, self.Total()))) {
			return
		}
		//spin
	}
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func New() ProgressBar {
	return &simple{}
}
