package main

type task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"due_date"`
	Status      string `json:"status"`
}

func (t task) validate() bool {
	return t.Title != "" && t.Description != "" && t.Status != ""
}

func (t task) validateTitle() bool {
	return t.Title != ""
}

func (t task) validateDescription() bool {
	return t.Description != ""
}

func (t task) validateStatus() bool {
	return t.Status != ""
}

func (t task) validateDueDate() bool {
	return t.DueDate != 0
}

func (t *task) updateTask(task task) {
	if task.validateTitle() {
		t.Title = task.Title
	}
	if task.validateDescription() {
		t.Description = task.Description
	}
	if task.validateStatus() {
		t.Status = task.Status
	}
	if task.validateDueDate() {
		t.DueDate = task.DueDate
	}
}
