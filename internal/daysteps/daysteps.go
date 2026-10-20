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
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("некорректный формат строки: ожидается 2 поля")
	}

	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("не удалось распарсить количество шагов: %w", err)
	}
	if ds.Steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("не удалось распарсить продолжительность прогулки: %w", err)
	}
	if ds.Duration <= 0 {
		return errors.New("продолжительность прогулки должна быть положительной")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("не удалось рассчитать калории при ходьбе: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
