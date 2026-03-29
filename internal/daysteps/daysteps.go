package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sliceStrings := strings.Split(datastring, ",")
	if len(sliceStrings) != 2 {
		return errors.New("incorrect count of data")
	}
	steps, err := strconv.Atoi(sliceStrings[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect 0 steps")
	}
	ds.Steps = steps
	timeWalking, err := time.ParseDuration(sliceStrings[1])
	if err != nil {
		return err
	}
	if timeWalking <= 0 {
		return errors.New("incorrect 0 time")
	}
	ds.Duration = timeWalking
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	d := spentenergy.Distance(ds.Steps, ds.Height)
	spCall, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	st := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, d, spCall)
	return st, nil
}
