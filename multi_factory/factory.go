package multi_factory

import "reflect"

type IHumanFactory interface {
	CreateHuman(humanFac interface{}) Human
}

type HumanFactory struct{}

func (hc *HumanFactory) CreateHuman(humanFac interface{}) Human {
	// typeOfHuman := reflect.TypeOf(humanFac).Name()
	m := reflect.ValueOf(&BlackHumanFactory{}).MethodByName("CreateHuman")
	// m := reflect.ValueOf(&humanFac).MethodByName("CreateHuman")
	m.Call(nil)
	//	return CreateHuman(typeOfHuman)
	return new(BlackHuman)
}

type BlackHumanFactory struct{}
type YellowHumanFactory struct{}
type WhiteHumanFactory struct{}

func (b *BlackHumanFactory) CreateHuman() Human {
	return new(BlackHuman)
}

func (y *YellowHumanFactory) CreateHuman() Human {
	return new(YellowHuman)
}

func (y *WhiteHumanFactory) CreateHuman() Human {
	return new(WhiteHuman)
}

func (hc *HumanFactory) CreateHumanByType(human interface{}) Human {
	typeOfHuman := reflect.TypeOf(human).Name()

	return CreateHuman(typeOfHuman)
}
