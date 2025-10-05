package spentcalories

import (
	"errors" 
    "fmt"    
    "log"    
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
func checkParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return errors.New("все входные параметры (шаги, вес, рост, продолжительность) должны быть больше нуля")
	}
	return nil
}
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("invalid data format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps value: %w", err)
	}
	activity:= parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration value: %w", err)
	}
return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	speed:= distanceKm / duration.Hours()
	return speed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}
	
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes() 
	
	
	calories := (weight * speed * durationInMinutes) / minInH
	
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}
	
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes() 
	
	
	baseCalories := (weight * speed * durationInMinutes) / minInH
	
	
	calories := baseCalories * walkingCaloriesCoefficient
	
	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println("Ошибка парсинга тренировки:", err)
		return "", err
	}
	
	
	if err := checkParams(steps, weight, height, duration); err != nil {
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
		log.Println("Ошибка расчета калорий:", err)
		return "", err
	}

	
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity, durationHours, distanceKm, meanSpeed, calories), nil
}