package domain

type UserID struct {
	ID string
	User User
}

//NewUserID -
func NewUserID() *UserID {
	return &UserID{}
}
