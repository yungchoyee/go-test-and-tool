package leecode

type NewError struct {
	s string
}

func (e *NewError)NewError(s string){
	e.s = s
}

func (e *NewError)Error() string{
	return e.s
}
