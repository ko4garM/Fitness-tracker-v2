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
		err := DataParser.Parse(dp, v)
		if err != nil {
			log.Println(err)
			continue
		}
		st, err := DataParser.ActionInfo(dp)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(st)
	}

}
