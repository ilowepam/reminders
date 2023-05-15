package repository

import "reminders/pkg/model"

type InMemoryStorageImpl struct {
	reminders *[]model.Reminder
}

func (storage InMemoryStorageImpl) Get(id int) *model.Reminder {
	return &model.Reminder{1, "Implement me!", false}
}

func (storage InMemoryStorageImpl) GetAll() []model.Reminder {
	return []model.Reminder{
		{1, "Implement me!", false},
		{2, "Implement too!", true},
	}
}
