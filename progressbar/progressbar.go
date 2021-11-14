package progressbar

type ProgressBar interface {
	Total() uint
	Done() uint
	Progress() float32
	Set(done uint)
}

type normalBar struct {
	done uint
}

func (self *normalBar) Total() uint {
	return 100
}

func (self *normalBar) Done() uint {
	return self.done
}

func (self *normalBar) Progress() float32 {
	return float32(self.Done()) / float32(self.Total())
}

func (self *normalBar) Set(done uint) {
	self.done = min(done, self.Total())
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
