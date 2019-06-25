package domain

type User struct {
	Login, 
	Password,
	UserRights string
}

//NewUser return empty User's instance
func NewUser() User {
	return User{}
}

