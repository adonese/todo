package main

import (
	"errors"
	"time"
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
	ID int
	CreateAt time.Time
	Status bool
	Description string
	Due  time.Time // default time.Now() + add(24hrs)
}

func (s *Storage)isValid()bool{
	if s.Description != "" {
		return true
	}
	return false
}


type User struct {
	Username string
	ID int
	Tasks []Storage
	Users []User
}

func (u *User)checkTask(t Storage) (bool){
	return true
}

func (u *User)append(t Storage) (bool){
	if u.checkTask(t){
		u.Tasks = append(u.Tasks, t)
		return true
	}
	return false
}

func (u *User)getUser(user string) (*User, error){
	for _, v := range u.Users{
		if user == v.Username {
			return &v, nil
		}
	}
	return &User{}, userNotFound
}


func (u *User)GetUser(username string) (error){
	res, err := u.getUser(username)
	u = res
	return err
}

func (u *User)GetTasks(username string) ([]Storage, error){

	res, err := u.getUser(username)
	return res.Tasks, err

}

func (u *User)GetTaskByID(username string, id int) (Storage, error){
	res, err := u.getUser(username)
	if err != nil {
		return Storage{}, err
	}

	for _, v := range res.Tasks {
		if v.ID == id {
			return v, nil
		}
	}
	return Storage{}, genericError
}


func (u *User)NewTask(username string, s Storage) error{
	if s.isValid(){
		u.Tasks = append(u.Tasks, s)
		return nil
	}else{
		return genericError
	}
}

func (u *User)UpdateTask(username string, id int, t Storage) (*Storage, error) {
	u1, err := u.getUser(username)
	if err != nil {
		return nil, err
	}
	for _, v := range u1.Tasks {
		if v.ID == id {
			v = t
		}
	}
	return nil, userNotFound

}


func (u *User)DeleteTask(username string, id int) error{
	u1, err := u.getUser(username)
	if err != nil {
		return err
	}
	for k, v := range u1.Tasks {
		if v.ID == id {
			u1.Tasks =  append(u1.Tasks[:k], u1.Tasks[k+1:]...)
			return nil
		}
	}
	return  userNotFound

}


var (
	userNotFound = errors.New("user not found")
	genericError = errors.New("there is an error")
)