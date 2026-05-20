package task

import "time"

var tasks = make(map[int64]Task)
var oid int64 = 0

type Task struct {
	Id          int64
	Title       string
	Description string
	CreatedBy   string
	Deadline    time.Time
	Completed   bool
}

func New(title, description string, createdBy string, deadline time.Time, completed bool) Task {

	//Логика для пристота ID
	oid++
	return Task{
		Id:          oid,
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
		Deadline:    deadline,
		Completed:   completed,
	}
}

func (t *Task) Create() {
	tasks[t.Id] = *t
}

func (t *Task) Update() {
	tasks[t.Id] = *t
}

func Get(id int64) Task {
	return tasks[id]
}

func (t *Task) Delete() {
	delete(tasks, t.Id)
}
