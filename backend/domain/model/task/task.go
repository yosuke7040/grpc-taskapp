package task

import "time"

type Task struct {
	id          string
	userID      string
	name        string
	isCompleted bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(
	id string,
	userID string,
	name string,
	isCompleted bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Task {
	return &Task{
		id:          id,
		userID:      userID,
		name:        name,
		isCompleted: isCompleted,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) UserID() string {
	return t.userID
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) IsCompleted() bool {
	return t.isCompleted
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}
