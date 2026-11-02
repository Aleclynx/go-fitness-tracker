package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для парсинга и вывода информации
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info — обрабатываем все строки из dataset
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка получения информации:", err)
			continue
		}

		fmt.Println(info)
	}
}
