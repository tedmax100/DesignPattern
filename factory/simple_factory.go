package factory

import "fmt"

//翻译接口
type Translator interface {
	Translate(string) string
}

//德语翻译类
type GermanTranslator struct{}

func (*GermanTranslator) Translate(words string) string {

	return "德语"
}

//英语翻译类
type EnglishTranslator struct{}

func (*EnglishTranslator) Translate(words string) string {

	return "英语"
}

//日语翻译类
type JapaneseTranslator struct{}

func (*JapaneseTranslator) Translate(words string) string {

	return "日语"
}

func Create(lan int) Translator {
	var translator Translator

	//根据不同的语言种类，实例化不同的翻译类
	switch lan {
	case 1:
		translator = new(GermanTranslator)
	case 2:
		translator = new(EnglishTranslator)
	case 3:
		translator = new(JapaneseTranslator)
	default:
		panic("no such translator")
	}

	return translator
}

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
	case "Black":
		return new(BlackHuman)
	case "Yellow":
		return new(YellowHuman)
	case "White":
		return new(WhiteHuman)
	default:
		panic("no such human")
	}
}
