package types

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Remainders reminder[] `json:"remainders"`

}

func (u *User) ValidateUser() {

}
