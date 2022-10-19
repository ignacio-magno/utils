package period

import "strconv"

type PeriodFlashBack struct {
	Period
}

func NewPeriodFlashBack(month, year string) (*PeriodFlashBack, error) {
	period, err := NewPeriod(month, year)
	if err != nil {
		return &PeriodFlashBack{}, nil
	}

	return &PeriodFlashBack{
		Period: *period,
	}, nil
}

func NewPeriodFlashBackFromPeriod(period Period) *PeriodFlashBack {
	return &PeriodFlashBack{
		Period: period,
	}
}

func (p *PeriodFlashBack) ReturnOneMonth() *PeriodFlashBack {
	var (
		newMonth int
		newYear  int
	)
	if p.GetMonthInt()-1 == 0 {
		newMonth = 12
		newYear = p.Year - 1
	} else {
		newMonth = p.GetMonthInt() - 1
		newYear = p.Year
	}

	andYear, err := obtainIdPeriodFromMonthAndYear(monthIntToStringWithCeros(newMonth), strconv.Itoa(newYear))
	if err != nil {
		panic(err)
	}

	p.Id = andYear
	p.Month = monthIntToStringWithCeros(newMonth)
	p.Year = newYear

	return p
}
