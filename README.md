# Todo App
Todo app backend using Go

# API Documentation

### Task JSON Data
* **id**: The unique identifier of the task.
* **title**: The title of the task.
* **description**: The description of the task.
* **due_date**: The due date of the task.
* **status**: The status of the task (e.g., "incomplete", "completed").
* #### GET ``/tasks``

```
Returns all the tasks that are available
```
* #### GET ``/tasks/:id``
```
Returns the associated task with that id
```
* #### POST ``/tasks``
```
Adds a task to the list

Params:

id: The unique identifier of the task.
title: The title of the task.
description: The description of the task.
due_date: The due date of the task.
status: The status of the task (e.g., "incomplete").
```
* #### PUT ``/tasks/:id``
```
Updates the associated task

Params:

title: The title of the task.
description: The description of the task.
due_date: The due date of the task.
status: The status of the task (e.g., "incomplete").
```
* #### PATCH ``/tasks/:id``
```
Updates the associated task with optinial params

Params:

title: The title of the task. (optinal)
description: The description of the task. (optinal)
due_date: The due date of the task. (optinal)
status: The status of the task (e.g., "incomplete"). (optinal)
```
* #### DELETE ``/tasks/:id``
```
Deletes the associated task
```