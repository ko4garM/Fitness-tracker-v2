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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sliceStrings := strings.Split(datastring, ",")
	if len(sliceStrings) != 3 {
		return errors.New("incorrect count of data")
	}
	steps, err := strconv.Atoi(sliceStrings[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect 0 steps")
	}
	t.Steps = steps
	t.TrainingType = sliceStrings[1]
	timeWalking, err := time.ParseDuration(sliceStrings[2])
	if err != nil {
		return err
	}
	if timeWalking <= 0 {
		return errors.New("incorrect 0 time")
	}
	t.Duration = timeWalking
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	d := spentenergy.Distance(t.Steps, t.Height)
	mnSp := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		spCall, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		fStr := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType) +
			fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours()) +
			fmt.Sprintf("Дистанция: %.2f км.\n", d) +
			fmt.Sprintf("Скорость: %.2f км/ч\n", mnSp) +
			fmt.Sprintf("Сожгли калорий: %.2f\n", spCall)
		return fStr, nil
	case "Бег":
		spCall, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		fStr := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType) +
			fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours()) +
			fmt.Sprintf("Дистанция: %.2f км.\n", d) +
			fmt.Sprintf("Скорость: %.2f км/ч\n", mnSp) +
			fmt.Sprintf("Сожгли калорий: %.2f\n", spCall)
		return fStr, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
