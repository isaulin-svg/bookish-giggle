package daysteps

import (
	"errors" 
	"fmt"   
	"log"    
	"strconv" 
	"strings"
	"time"
	"BOOKISH-GIGGLE/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid data format ")
}
steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps value: %w", err)
	}
	if steps <= 0 {	
		return 0, 0, errors.New("steps must be a positive integer")
	}
	duration, err := time.ParseDuration(parts[1])
		if err != nil {
		return 0, 0, fmt.Errorf("invalid duration value: %w", err)
	}
	return steps, duration, nil
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Error parsing data:", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("Error calculating calories:", err)
		return ""
	}
	return fmt.Sprintf("Вы прошли %d шагов за %s (%.2f км) и сожгли %.2f ккал.",  steps, distanceKm, calories)
}

