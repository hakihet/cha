package cha

type Cha struct {
	quit chan struct{}
	done chan struct{}
}

func (cha Cha) Stop() {
	close(cha.quit)
}

func (cha Cha) Quit() bool {
	select {
	case <-cha.quit:
		return true
	default:
		return false
	}
}

func (cha Cha) Done() {
	close(cha.done)
}

func (cha Cha) Wait() {
	<-cha.done
}

func New() Cha {
	return Cha{make(chan struct{}), make(chan struct{})}
}
