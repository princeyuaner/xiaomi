package main

import (
	"mi/Function"
)

func main() {
	//xiaomi.LogInit()
	xiaomi.Login()
	//xiaomi.AHL()
	//xiaomi.HuoDongLogin()
	xiaomi.Remind()

	for {
		xiaomi.BuyGood()
	}
}
