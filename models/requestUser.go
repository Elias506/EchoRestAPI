package models

type RequestUser struct {
	Name string `json:"name"`
}

func (u *User) NewUser(rU *RequestUser) {
	// without ID
	u.Name = rU.Name
	//...
}