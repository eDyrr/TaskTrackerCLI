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


**building the project**
1. *init*ing the project by `exec` the command `go mod init <module-name>`
2. **defining the command**:
    to define the command that I need to type into my CLI inorder for the app to get called we need a `Go` package named `flag`, which provides a simple way to define command line flags and arguments.
    in our case 


**the algorithm**:
1. writing the input
2. reading the input
3. analysing the input
4. calling the adequate function


**how the command will look like:**
disclaimer: this is just the first version of the command
./main -task=add "go meet joe at 7 pm"
./main -task=list -filter=done