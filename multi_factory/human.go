package multi_factory

import "fmt"

type Human interface {
	GetColor() string
	Talk()
}

type BlackHuman struct{}

func (h *BlackHuman) GetColor() string {
	return "黑人膚色是黑色的。"
}

func (h *BlackHuman) Talk() {
	fmt.Println("黑人會說話，一般人聽不懂")
}

type YellowHuman struct{}

func (h *YellowHuman) GetColor() string {
	return "黃種人膚色是黃色的。"
}

func (h *YellowHuman) Talk() {
	fmt.Println("黃種人會說話，說的都是2byte的文字")
}

type WhiteHuman struct{}

func (h *WhiteHuman) GetColor() string {
	return "白人膚色是白色的。"
}

func (h *WhiteHuman) Talk() {
	fmt.Println("白人會說話，說的都是1byte的文字")
}

func CreateHuman(humanType string) Human {
	switch humanType {
	case "Black", "BlackHuman":
		return new(BlackHuman)
	case "Yellow", "YellowHuman":
		return new(YellowHuman)
	case "White", "WhiteHuman":
		return new(WhiteHuman)
	default:
		panic("no such human")
	}
}
