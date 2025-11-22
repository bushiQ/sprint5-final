package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("splitting error")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("steps ivalid format: %w", err)
	}
	if steps <= 0 {
		return errors.New("incorrect steps")
	}
	t.Steps = steps
	parsedTime, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("time parsing error: %w", err)
	}
	if parsedTime <= 0 {
		return errors.New("incorrect time")
	}
	t.Duration = parsedTime
	t.TrainingType = parts[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 {
		return "", errors.New("incorrect steps")
	}
	if t.Duration <= 0 {
		return "", errors.New("incorrect time")
	}
	dist := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		spentCals, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		inf := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, meanSpeed, spentCals)
		return inf, err
	case "Бег":
		spentCals, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		inf := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, meanSpeed, spentCals)
		return inf, err
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
