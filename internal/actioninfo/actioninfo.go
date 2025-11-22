package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
		}
		inf, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(inf)
	}
}
