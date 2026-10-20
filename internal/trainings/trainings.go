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

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("не удалось распарсить количество шагов: %w", err)
	}
	if t.Steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("не удалось распарсить продолжительность тренировки: %w", err)
	}
	if t.Duration <= 0 {
		return errors.New("продолжительность тренировки должна быть положительной")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64

	switch t.TrainingType {
	case "Бег":
		c, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("не удалось рассчитать калории при беге: %w", err)
		}
		calories = c
	case "Ходьба":
		c, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("не удалось рассчитать калории при ходьбе: %w", err)
		}
		calories = c
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories,
	), nil
}
