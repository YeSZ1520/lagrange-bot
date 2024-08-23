package model

type Aborts struct {
	abort bool
}

func (a *Aborts) Abort() {
	a.abort = true
}

func (a *Aborts) IsAbort() bool {
	return a.abort
}
