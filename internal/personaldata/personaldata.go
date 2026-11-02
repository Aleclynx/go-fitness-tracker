package personaldata

import "fmt"

// Personal — структура для данных юзера
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print — выводим информацию о пользователе
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
