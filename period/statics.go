package period

import (
	"strconv"
	"time"
)

// GetDescriptionToString yy-mmmm
func GetDescriptionToString(idPeriod string) (string, error) {
	fromId, err := NewPeriodFromId(idPeriod)
	if err != nil {
		return "", err
	}
	return fromId.GetPeriodString(), nil
}

// GetPeriodFromIdStringLocal julio 2022
func GetDescriptionToStringLocal(idPeriod string) (string, error) {
	fromId, err := NewPeriodFromId(idPeriod)
	if err != nil {
		return "", err
	}
	return fromId.GetPeriodStringLocal(), nil
}

func NewFromCurrentPeriod() *Period {
	date := time.Now()
	month := int(date.Month())
	year := date.Year()

	period, err := NewPeriod(monthIntToStringWithCeros(month), strconv.Itoa(year))
	if err != nil {
		panic(err)
	}
	return period

}
