Tasks bot design

keyword task (configurable)
description
madeby
assigned_to
due (default 24hr)

Notify(task) []Users

RespondToTask(ID): (
Completed
IsPending
Cancelled)

SearchTasksToMark(query) ID

Input feed:

- http calls
- telegram groups

Permissions:

- send DMs
/remindme to drink water at two hours from now
/remindme to send to abbas rummly at two hours from now
/remindme to 


Specs:

- ability to send messages
- ability to be called in telegram groups
- ability to be called in telegram chat
- ability to store context (without DB)
- ability to schedule calls


System

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
