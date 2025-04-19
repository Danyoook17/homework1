package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)
var ErrInvalidDataFormat = errors.New("invalid data format ")

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	substrings := strings.Split(data, ",")

	if len(substrings) != 2 {
		return 0, 0, ErrInvalidDataFormat
	}

	stepsNumber, err := strconv.Atoi(substrings[0])
	if err != nil {
		return 0, 0, ErrInvalidDataFormat
	}

	if stepsNumber <= 0 {
		return 0, 0, fmt.Errorf("negative number of steps: %w", ErrInvalidDataFormat)
	}

	

	trainingDuration, err := time.ParseDuration(substrings[1])
	if err != nil {
		return 0, 0, errors.Join(err, ErrInvalidDataFormat)
	}

	if trainingDuration <= 0 {
		return 0, 0, fmt.Errorf("negative training duration: %w", ErrInvalidDataFormat)
	}

	return stepsNumber, trainingDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	return ""
}
