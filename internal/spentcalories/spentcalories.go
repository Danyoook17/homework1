package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var ErrInvalidDataFormat = errors.New("invalid data format")

func parseTraining(data string) (int, string, time.Duration, error) {
	substrings := strings.Split(data, ",")
	if len(substrings) != 3 {
		return 0, "", 0, fmt.Errorf("Invalid data format: %w", ErrInvalidDataFormat)
	}

	stepsNumber, err := strconv.Atoi(substrings[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid data format: %w", ErrInvalidDataFormat)
	}

	if stepsNumber <= 0 {
		return 0, "", 0, fmt.Errorf("negative number of steps: %w", ErrInvalidDataFormat)
	}

	trainingDuration, err := time.ParseDuration(substrings[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid data format: %w", ErrInvalidDataFormat)
	}

	if trainingDuration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid data format: %w", ErrInvalidDataFormat)
	}

	return stepsNumber, substrings[1], trainingDuration, nil
}

func distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	distanceInMetres := stepLength * float64(steps)

	distanceInKiMetres := distanceInMetres / mInKm

	return distanceInKiMetres
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)

	hours := duration.Hours()

	avgSpeed := dist / hours

	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	stepsNumber, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var calories float64

	switch activity {

	case "Бег":
		calories, err = RunningSpentCalories(stepsNumber, weight, height, duration)
		if err != nil {
			return "", err
		}

	case "Ходьба":
		calories, err = WalkingSpentCalories(stepsNumber, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	distance := meanSpeed(stepsNumber, height, duration) * duration.Hours()

	speed := meanSpeed(stepsNumber, height, duration)

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), distance, speed, calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("steps can not be negative")
	}

	if weight <= 0 {
		return 0, errors.New("weight can not be negative")
	}

	if height <= 0 {
		return 0, errors.New("height can not be negative")
	}

	if duration <= 0 {
		return 0, errors.New("duration can not be negative")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("steps can not be negative")
	}

	if weight <= 0 {
		return 0, errors.New("weight can not be negative")
	}

	if height <= 0 {
		return 0, errors.New("height can not be negative")
	}

	if duration <= 0 {
		return 0, errors.New("duration can not be negative")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	walkingCalories := calories * walkingCaloriesCoefficient

	return walkingCalories, nil
}
