package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return fmt.Errorf("incorrect data format: %s", datastring)
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps is set incorrectly: %d", steps)
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return fmt.Errorf("time is set incorrectly: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("time is set incorrectly: %w", err)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("incorrect calories counting: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
