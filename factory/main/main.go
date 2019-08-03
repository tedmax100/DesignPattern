package main

import (
	"fmt"
	"time"

	"github.com/tedmax100/DesignPattern/factory"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)
	}()

	human := factory.CreateHuman("Yellow")
	human.Talk()
	fmt.Println(human.GetColor())

	humanFactory := new(factory.HumanFactory)
	human2 := humanFactory.CreateHumanByType(factory.BlackHuman{})

	human2.Talk()
	fmt.Println(human2.GetColor())

	/* 	a := factory.BlackHuman{}
	   	bb := reflect.TypeOf(a)
	   	fmt.Println(bb.Name()) */
}
