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
