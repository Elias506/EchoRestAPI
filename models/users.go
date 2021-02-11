package models

type Users struct {
	Users   []User `json:"users"`
	Counter int    `json:"counter"`
}
