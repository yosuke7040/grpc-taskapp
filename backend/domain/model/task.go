package model

import "time"

// エンティティなので、IDは区別するために普遍かつ一意の必要があると思うが、
// そのほかのフィールドは変更される場合があるのでパブリックにしてみた
type Task struct {
	id          ID
	UserID      ID
	Name        string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(
	id ID,
	userID ID,
	name string,
	isCompleted bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Task {
	return &Task{
		id:          id,
		UserID:      userID,
		Name:        name,
		IsCompleted: isCompleted,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (t *Task) ID() ID {
	return t.id
}

func (t *Task) ChangeName(name string) {
	t.Name = name
}

func (t *Task) Complete() {
	t.IsCompleted = true
}

func (t *Task) Uncomplete() {
	t.IsCompleted = false
}

func (t *Task) UpdateUpdatedAt(time time.Time) {
	t.UpdatedAt = time
}
