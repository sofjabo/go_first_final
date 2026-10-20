package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
    Parse(datastring string) error
    ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, s := range dataset {
		if err := dp.Parse(s); err != nil {
			log.Printf("ошибка парсинга %q: %v", s, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
