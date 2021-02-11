package repository

import (
	"encoding/json"
	"github.com/elias506/EchoRestAPI/models"
	"io/ioutil"
	"os"
	"github.com/elias506/EchoRestAPI/repository/models"
)

func (db FileDB) open() error {
	var err error
	db.file, err = os.Open(db.path)
	if err != nil {
		return err
	}
	return nil
}

func (db *FileDB) close() {
	_ = db.file.Close()
}

func (db *FileDB) read() (*Users, error) {
	body, err := ioutil.ReadAll(db.file)
	if err != nil {
		return nil, err
	}
	u, err := unmarshal(body)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *FileDB) write(users *Users) error {
	newBody, err := marshal(users)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(db.path, newBody, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (db *FileDB) GetUser(id int) (*User, error) {
	err := db.open()
	if err != nil {
		return nil, err
	}
	defer db.close()
	users, err := db.read()
	if err != nil {
		return nil, err
	}

	// ...
	for _, u := range users.Users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, nil
}

func (db *FileDB) GetUsers() (*[]User, error) {
	err := db.open()
	if err != nil {
		return nil, err
	}
	defer db.close()
	users, err := db.read()
	if err != nil {
		return nil, err
	}

	return &users.Users, nil
}

func (db *FileDB) DeleteUser(id int) (int, error) {
	err := db.open()
	if err != nil {
		return -1, err
	}
	defer db.close()
	users, err := db.read()
	if err != nil {
		return -1, err
	}

	for i, u := range users.Users {
		if u.ID == id {
			users.Users = remove(users.Users, i)
			err := db.write(users)
			if err != nil {
				return -1, nil
			}
			return id, nil
		}
	}
	return -1, nil
}

func (db *FileDB) UpdateUser(id int, reqUser *RequestUser) (int, error) {
	err := db.open()
	if err != nil {
		return -1, err
	}
	defer db.close()
	users, err := db.read()
	if err != nil {
		return -1, err
	}
	
	for i, u := range users.Users {
		if id == u.ID {

			users.Users[i].NewUser(reqUser)
			err := db.write(users)
			if err != nil {
				return -1, err
			}
			return id, nil
		}
	}
	return -1, nil
}

func (db *FileDB) AddUser(reqUser *RequestUser) (int, error) {
	err := db.open()
	if err != nil {
		return -1, err
	}
	defer db.close()
	users, err := db.read()
	if err != nil {
		return -1, err
	}
	user := &User{
		ID: users.Counter,
	}
	users.Counter++
	user.NewUser(reqUser)
	users.Users = append(users.Users, *user)

	err = db.write(users)
	if err != nil {
		return -1, nil
	}

	return user.ID, nil
}

func unmarshal(body []byte) (*Users, error) {
	var u Users
	err := json.Unmarshal(body, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func marshal(users *Users) ([]byte, error) {
	body, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func remove(slice []User, i int) []User {
	return append(slice[:i], slice[i+1:]...)
}
