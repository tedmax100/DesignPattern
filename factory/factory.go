package factory

type IHumanFactory interface {
	CreateHuman(humanType string) Human
}

type HumanFactory struct{}

func (hc *HumanFactory) CreateHuman(humanType string) Human {
	return CreateHuman(humanType)
}
