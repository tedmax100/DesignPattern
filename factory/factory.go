package factory

import "reflect"

type IHumanFactory interface {
	CreateHuman(humanType string) Human
}

type HumanFactory struct{}

func (hc *HumanFactory) CreateHuman(humanType string) Human {
	return CreateHuman(humanType)
}

func (hc *HumanFactory) CreateHumanByType(human interface{}) Human {
	typeOfHuman := reflect.TypeOf(human).Name()

	return CreateHuman(typeOfHuman)
}
