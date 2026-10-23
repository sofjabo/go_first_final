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
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("некорректный формат строки: ожидается 3 поля")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("не удалось распарсить количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("не удалось распарсить продолжительность тренировки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность тренировки должна быть положительной")
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("не удалось рассчитать калории: %w", err)
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories,
	), nil
}
