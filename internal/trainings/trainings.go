package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training — данные о тренировке
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse — разбираем строку с тренировкой
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки тренировки")
	}

	// убираем пробелы если есть
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return errors.New("некорректное количество шагов")
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil || duration <= 0 {
		return errors.New("некорректная продолжительность тренировки")
	}
	t.Duration = duration

	return nil
}

// ActionInfo — собираем строку с информацией о тренировке
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return info, nil
}
