package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс, который реализуют структуры Training и DaySteps.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info принимает слайс строк и объект, реализующий интерфейс DataParser.
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