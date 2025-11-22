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

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("steps must be above zero")
	case weight <= 0:
		return 0, errors.New("weight is incorrect")
	case height <= 0:
		return 0, errors.New("height is incorrect")
	case duration <= 0:
		return 0, errors.New("duration is incorrect")
	default:
		meanSpeed := MeanSpeed(steps, height, duration)
		durationInMinutes := duration.Minutes()
		return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("steps must be above zero")
	case weight <= 0:
		return 0, errors.New("weight is incorrect")
	case height <= 0:
		return 0, errors.New("height is incorrect")
	case duration <= 0:
		return 0, errors.New("duration is incorrect")
	default:
		meanSpeed := MeanSpeed(steps, height, duration)
		durationInMinutes := duration.Minutes()
		return (weight * meanSpeed * durationInMinutes) / minInH, nil
	}
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceM := stepLength * float64(steps)
	return distanceM / mInKm
}
