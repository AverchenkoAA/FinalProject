package domain

type PC struct {
	ID string
	Computer Computer
}

//NewPC -
func NewPC() *PC {
	return &PC{}
}
