package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories вычисляет количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем, что все параметры имеют смысл.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры для расчета калорий при ходьбе")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	// Базовая формула калорий.
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// Для ходьбы применяем корректирующий коэффициент.
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories вычисляет количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры для расчета калорий при беге")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Если продолжительность не положительная или шаги отрицательные — скорость 0.
	if duration <= 0 || steps < 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

// Distance вычисляет дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	// Длина шага зависит от роста.
	stepLength := height * stepLengthCoefficient
	// Общая длина в метрах.
	distanceInMeters := float64(steps) * stepLength
	// Переводим в километры.
	return distanceInMeters / mInKm
}