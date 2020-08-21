package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

const (
	notFinished = iota + 1
	isPending
	delayed
	done
)

/*

Storage

ID
CreateAt
Status
Description
Due default time.Now() + add(24hrs)

Methods

- NewTask() error
- GetByID() Storage, error
- UpdateByID(Storage) error
- Delete(ID) error

User

Username
ID
Tasks []Storage

Methods

- GetUser(telegram_username) User, error
- GetTasks(telegram_username) []Storage, error
- GetTaskByID(telegram_username, id) Storage, error
- UpdateTask(telegram_username, id) error
- DeleteTask(telegram_username, id) error
*/

//Storage building block for a task
type Storage struct {
	ID          int
	CreateAt    time.Time
	Status      int
	Description string
	Due         time.Time // default time.Now() + add(24hrs)
	IsDone      bool
}

//Message contains telegram content
type Message struct {
	message string
}

func (s *Storage) isValid() bool {
	return s.Description == ""
}

//User contains storage for all users data
type User struct {
	Username string
	ID       int
	Tasks    []Storage
	Users    []User
}

func (u *User) checkTask(t Storage) bool {
	return true
}

func (u *User) append(t Storage) bool {
	if u.checkTask(t) {
		u.Tasks = append(u.Tasks, t)
		return true
	}
	return false
}

func (u *User) getUser(user string) (*User, error) {
	for _, v := range u.Users {
		if user == v.Username {
			return &v, nil
		}
	}
	return &User{}, errNotFound
}

func (u *User) GetUser(username string) error {
	res, err := u.getUser(username)
	u = res
	return err
}

func (u *User) GetTasks(username string) ([]Storage, error) {

	res, err := u.getUser(username)
	return res.Tasks, err

}

func (u *User) GetTasksString(username string) ([]string, error) {

	_, err := u.getUser(username)
	return []string{}, err

}

func (u *User) GetTaskByID(username string, id int) (Storage, error) {
	res, err := u.getUser(username)
	if err != nil {
		return Storage{}, err
	}

	for _, v := range res.Tasks {
		if v.ID == id {
			return v, nil
		}
	}
	return Storage{}, errGeneric
}

func (u *User) NewTask(username string, s Storage) error {
	if s.isValid() {
		u.Tasks = append(u.Tasks, s)
		return nil
	} else {
		return errGeneric
	}
}

func (u *User) UpdateTask(username string, id int, t Storage) (*Storage, error) {
	u1, err := u.getUser(username)
	if err != nil {
		return nil, err
	}
	for _, v := range u1.Tasks {
		if v.ID == id {
			if t.Due.IsZero() {
				t.Due = v.Due
			}
			if t.Description == "" {
				t.Description = v.Description
			}
			// fix is done
			v = t
		}
	}
	return nil, errNotFound

}

func (u *User) DeleteTask(username string, id int) error {
	u1, err := u.getUser(username)
	if err != nil {
		return err
	}
	for k, v := range u1.Tasks {
		if v.ID == id {
			u1.Tasks = append(u1.Tasks[:k], u1.Tasks[k+1:]...)
			return nil
		}
	}
	return errNotFound

}

// This example transmits a value that implements the custom encoding and decoding methods.
func encodeDecode() error {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Storage{Description: "my work"})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data.dump", network.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func readFile() (Storage, error) {
	// Create a decoder and receive a value.
	net, err := os.Open("data.dump")
	if err != nil {
		return Storage{}, err
	}

	dec := gob.NewDecoder(net)
	var v Storage
	err = dec.Decode(&v)
	if err != nil {
		return Storage{}, err
	}
	return v, nil

}

var (
	errNotFound = errors.New("user not found")
	errGeneric  = errors.New("there is an error")
)
