package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, errors.New("количество шагов должно быть больше нуля")
	case weight <= 0:
		return 0, errors.New("вес должен быть больше нуля")
	case height <= 0:
		return 0, errors.New("рост должен быть больше нуля")
	case duration <= 0:
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := spentCalories(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	return spentCalories(steps, weight, height, duration)
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
