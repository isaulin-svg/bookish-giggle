package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	mInKm = 1000.0 // количество метров в километре.
	minInH = 60.0  // количество минут в часе.
	stepLengthCoefficient = 0.45 

	runningCaloriesK1 = 18.0 // Коэффициент скорости для бега
	runningCaloriesK2 = 20.0 // Базовый коэффициент для бега
	
	walkingCaloriesK1 = 0.35 // Коэффициент скорости для ходьбы
	walkingCaloriesK2 = 1.8  // Базовый коэффициент для ходьбы
	walkingCaloriesCoefficient = 0.02
)



func checkParams(weight, height float64) error {
	if weight <= 0 || height <= 0 {
		return errors.New("вес и рост должны быть больше нуля")
	}
	return nil
}
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("invalid data format: expected steps,activity,duration")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps value: %w", err)
	}

	
	if steps <= 0 {
		return 0, "", 0, errors.New("шаги должны быть положительным числом")
	}

	activity := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration value: %w", err)
	}

	
	if duration <= 0 {
		return 0, "", 0, errors.New("продолжительность должна быть больше нуля")
	}
	return steps, activity, duration, nil 
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	speed := distanceKm / duration.Hours()
	return speed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if err := checkParams(weight, height); err != nil {
		return 0, err
	}
	if steps <= 0 {
		return 0, errors.New("шаги должны быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes() 
	calories := ((runningCaloriesK1 * speed) + runningCaloriesK2) * weight * durationInMinutes / minInH
	
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if err := checkParams(weight, height); err != nil {
		return 0, err
	}
	if steps <= 0 {
		return 0, errors.New("шаги должны быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}
	
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes() 
	baseRunningCalories := ((runningCaloriesK1 * speed) + runningCaloriesK2) * weight * durationInMinutes / minInH
	calories := baseRunningCalories / 2.0 
	
	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println("Ошибка парсинга тренировки:", err)
		return "", err
	}
	
	if err := checkParams(weight, height); err != nil {
		log.Println("Ошибка входных параметров:", err)
		return "", err
	}

	distanceKm := distance(steps, height)
	meanSpeed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	var calories float64
	
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	
	if err != nil {
		
		return "", err
	}

	
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, durationHours, distanceKm, meanSpeed, calories), nil
}