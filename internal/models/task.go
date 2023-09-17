package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"due_date"`
	Status      string `json:"status"`
}

func (t *Task) Validate() bool {
	return t.Title != "" && t.Description != "" && t.Status != ""
}

func (t *Task) ValidateTitle() bool {
	return t.Title != ""
}

func (t *Task) ValidateDescription() bool {
	return t.Description != ""
}

func (t *Task) ValidateStatus() bool {
	return t.Status != ""
}

func (t *Task) ValidateDueDate() bool {
	return t.DueDate != 0
}

func (t *Task) UpdateTask(task Task) {
	if task.ValidateTitle() {
		t.Title = task.Title
	}
	if task.ValidateDescription() {
		t.Description = task.Description
	}
	if task.ValidateStatus() {
		t.Status = task.Status
	}
	if task.ValidateDueDate() {
		t.DueDate = task.DueDate
	}
}
