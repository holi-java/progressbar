package progressbar

import . "sync/atomic"

type ProgressBar interface {
	Total() uint
	Done() uint
	Progress() float32
	Set(done uint)
}

type normalBar struct {
	done uint32
}

func (self *normalBar) Total() uint {
	return 100
}

func (self *normalBar) Done() uint {
	return uint(LoadUint32(&self.done))
}

func (self *normalBar) Progress() float32 {
	return float32(self.Done()) / float32(self.Total())
}

func (self *normalBar) Set(done uint) {
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
	return &normalBar{}
}
