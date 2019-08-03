package main

import (
	"fmt"
	"time"

	factory "github.com/tedmax100/DesignPattern/multi_factory"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)
	}()

	facotry := new(factory.HumanFactory)
	human := facotry.CreateHuman(factory.BlackHumanFactory{})
	human.Talk()
	fmt.Println(human.GetColor())
}
