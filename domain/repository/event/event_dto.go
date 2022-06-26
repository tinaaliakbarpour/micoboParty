package event

import (
	"micobianParty/domain/entity"
)

var Repository EventRepository = &event{}

type EventRepository interface {
	GetAll() ([]entity.Event, error)
	Get(id uint) (*entity.Event, error)
}

type event struct{}
