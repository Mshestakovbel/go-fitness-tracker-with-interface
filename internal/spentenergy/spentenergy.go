package spentenergy

import (
	"fmt"
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
	if steps <= 0 {
		return 0, fmt.Errorf("steps are set incorrectly: %d", steps)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("time are set incorrectly: %v", duration)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight are set incorrectly: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height are set incorrectly: %.2f", height)
	}
	spd := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * spd * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps are set incorrectly: %d", steps)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("time are set incorrectly: %v", duration)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight are set incorrectly: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height are set incorrectly: %.2f", height)
	}
	spd := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * spd * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	d := Distance(steps, height)
	speed := d / duration.Hours()
	return speed
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}
