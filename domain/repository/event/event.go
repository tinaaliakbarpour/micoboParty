package event

import (
	"micobianParty/client/postgres"
	"micobianParty/domain/entity"
)

//GetAll will return all the events in db
func (event) GetAll() ([]entity.Event, error) {
	var events []entity.Event
	db := postgres.Storage.DB()
	if err := db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

//Get will query an event with specified id
func (event) Get(id uint) (*entity.Event, error) {
	var event entity.Event
	db := postgres.Storage.DB()
	if err := db.Where("id = ?", id).Find(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
