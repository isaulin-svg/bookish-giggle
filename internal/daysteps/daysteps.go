package daysteps

import (
	"errors" 
	"fmt"   
	"log"    
	"strconv" 
	"strings"
	"time"
	
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	
	stepLength = 0.65
	
	mInKm = 1000.0
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid data format: expected steps,duration")
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
	
	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive")
	}

	return steps, duration, nil
} 

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Error parsing data:", err)
		return ""
	}

	
	if weight <= 0 || height <= 0 {
		log.Println("Weight and height must be positive")
		return ""
	}
	
	
	distanceMeters := float64(steps) * stepLength
	
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("Error calculating calories:", err)
		return ""
	}

	
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", 
		steps, distanceKm, calories) 
}