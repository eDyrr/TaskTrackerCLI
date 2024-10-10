### Building a Task Tracker CLI tool

###### The requirements

- add, update and delete tasks.
- mark a task as in progress or done.
- list all tasks.
- list all tasks that are done.
- list all tasks that are not done.
- list all tasks that are in progress.

###### The design

now that the requirements are well defined, now I am about to design the app.

**the functions:**
- addition
- updation
- deletion
- update status
- list all
- list with filter

**the objects:** task, with a status attribute that has 3 values:
1. to do
2. done
3. in-progress

and the properties of the task are:
- `id`
- `description`
- `status`
- `createdAt`
- `updatedAt`

**the constraints:**
- the use of positional arugements in command line to accept user inputs.
- the use of a JSON file to store the tasks in the current directory.
- the JSON file should be created if it does not exist.
