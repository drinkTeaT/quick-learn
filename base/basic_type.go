package main

func tigerZoo() bool {
	taiGe := TaiGe{}
	taiGe.NAME = "tai ge"
	taiGe.JUMP = true
	taiGe.FOOD = "meat"
	taiGe.MALE = true
	taiGe.YEARS = 3
	taiGe.LENGTH = 2.1
	taiGe.ToyMap["ball"] = 3
	taiGe.ToyMap["shoe"] = 5
	return true
}

func main() {
	tigerZoo()
}

type Tiger struct {
	JUMP bool
	FOOD string
	NAME string
}

type TaiGe struct {
	Tiger
	YEARS       int
	LENGTH      float32
	NAME        string
	MALE        bool
	GoodFriends [2]string
	RECIPE      [...]string
	ToyMap      map[string]int
}
