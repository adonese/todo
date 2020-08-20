package main

import "time"

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

func (s *Storage)NewTask() error{
	return nil
}

func (s *Storage)GetByID(ID int) ( error){
	return nil
}

func (s *Storage)UpdateByID(ID int) error{
	return nil
}


func (s *Storage)Delete(ID int) error{
return nil
}



type User struct {
	Username string
	ID int
	Tasks []Storage
}

func (u *User)GetUser(username string) (error){
	return nil
}

func (u *User)GetTasks(username string) ([]Storage, error){
	return []Storage{}, nil
}
func (u *User)GetTaskByID(username string, id int) (Storage, error){
	return Storage{}, nil
}

func (u *User)UpdateTask(username string, id int) error{
	return nil
}

func (u *User)DeleteTask(username string, id int) error{
	return nil
}
