package repository

import (
	. "github.com/elias506/EchoRestAPI/models"
)

type IUserDB interface {
	GetUser(id int) (*User, error)
	GetUsers() (*[]User, error)
	DeleteUser(id int) (int, error)
	UpdateUser(id int, reqUser *RequestUser) (int, error)
	AddUser(reqUser *RequestUser) (int, error)
}

func IGetUser(i IUserDB, id int) (*User, error) {
	return i.GetUser(id)
}

func IGetUsers(i IUserDB) (*[]User, error) {
	return i.GetUsers()
}

func IDeleteUser(i IUserDB, id int) (int, error) {
	return i.DeleteUser(id)
}

func IUpdateUser(i IUserDB, id int, reqUser *RequestUser) (int, error) {
	return i.UpdateUser(id, reqUser)
}

func IAddUser(i IUserDB, reqUser *RequestUser) (int, error) {
	return i.AddUser(reqUser)
}
