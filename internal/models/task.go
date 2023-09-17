package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"due_date"`
	Status      string `json:"status"`
}

func (t Task) Validate() bool {
	return t.Title != "" && t.Description != "" && t.Status != ""
}

func (t Task) validateTitle() bool {
	return t.Title != ""
}

func (t Task) validateDescription() bool {
	return t.Description != ""
}

func (t Task) validateStatus() bool {
	return t.Status != ""
}

func (t Task) validateDueDate() bool {
	return t.DueDate != 0
}

func (t *Task) UpdateTask(task Task) {
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
