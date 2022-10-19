package period

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Implementa StorageMongo
type Period struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Month string             `bson:"month" json:"month"`
	Year  int                `bson:"year" json:"year"`
}

// NewPeriod Construct un período con el mes y año ingresado mm , yyyyy
func NewPeriod(month string, year string) (*Period, error) {
	yn, err := strconv.Atoi(year)
	if err != nil {
		return &Period{}, err
	}

	primitiveId, err := obtainIdPeriodFromMonthAndYear(month, year)

	return &Period{
		Id:    primitiveId,
		Month: month,
		Year:  yn,
	}, err
}

func NewPeriodFromInt(month int, year int) (*Period, error) {
	return NewPeriod(monthIntToStringWithCeros(month), strconv.Itoa(year))
}

func NewPeriodFromId(idString string) (*Period, error) {
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return &Period{}, err
	}
	for k, v := range MapPeriods {
		if v == id.Hex() {
			month := strings.Split(k, "-")[0]
			year := strings.Split(k, "-")[1]
			yearInt, err := strconv.Atoi(year)
			if err != nil {
				return &Period{}, err
			}
			return &Period{
				Id:    id,
				Month: month,
				Year:  yearInt,
			}, nil
		}
	}
	return &Period{}, fmt.Errorf("no se encontró el periodo")
}

// NewPeriodForPreviousPeriod Crea un Period basado en la fecha actual de la que se ejecuta el programa. con un mes anterior
func NewPeriodForPreviousPeriod() (*Period, error) {
	date := time.Now()
	month := int(date.Month()) - 1
	year := date.Year()
	if month == 0 {
		year = date.Year() - 1
		month = 12
	}

	return NewPeriod(monthIntToStringWithCeros(month), strconv.Itoa(year))
}

// GetPeriodString return in format mm-yyyy
func (p *Period) GetPeriodString() string {
	return fmt.Sprintf("%s-%d", p.Month, p.Year)
}

// GetPeriodStringLocal return in format local, ej: Enero 2020
func (p *Period) GetPeriodStringLocal() string {
	return fmt.Sprintf("%s %d", monthToLocal(p.Month), p.Year)
}

func (p *Period) GetMonthInt() int {
	return monthToInt(p.Month)
}

// pesca el actual periodo ya creado, y basado en esto, retrocede un mes, construyendo el objeto con todos los nuevos datos
func (p *Period) GeneratePeriodFlashBack() *PeriodFlashBack {
	return NewPeriodFlashBackFromPeriod(*p)
}

// NewPeriodConstructorWithString Construye un periodo de accuser a un string
// debe recibir un string de la siguiente forma aaaa-mm
func NewPeriodConstructorWithString(d string) (*Period, error) {
	month_year := strings.Split(d, "-")

	month := month_year[1]
	year, err := strconv.Atoi(month_year[0])

	if err != nil {
		return &Period{}, err
	}

	return NewPeriod(month, strconv.Itoa(year))
}

// Función implementada por storage mongo
func (p *Period) GetNameCollection() string {
	return "periods"
}

// return the id Period from the month and year
func obtainIdPeriodFromMonthAndYear(month, year string) (primitive.ObjectID, error) {
	id := MapPeriods[month+"-"+year]
	return primitive.ObjectIDFromHex(id)
}

func monthIntToStringWithCeros(i int) string {
	switch i {
	case 1:
		return "01"
	case 2:
		return "02"
	case 3:
		return "03"
	case 4:
		return "04"
	case 5:
		return "05"
	case 6:
		return "06"
	case 7:
		return "07"
	case 8:
		return "08"
	case 9:
		return "09"
	case 10:
		return "10"
	case 11:
		return "11"
	case 12:
		return "12"
	default:
		return "00"
	}
}

func monthToLocal(month string) string {
	switch month {
	case "01":
		return "Enero"
	case "02":
		return "Febrero"
	case "03":
		return "Marzo"
	case "04":
		return "Abril"
	case "05":
		return "Mayo"
	case "06":
		return "Junio"
	case "07":
		return "Julio"
	case "08":
		return "Agosto"
	case "09":
		return "Septiembre"
	case "10":
		return "Octubre"
	case "11":
		return "Noviembre"
	case "12":
		return "Diciembre"
	default:
		return "error"
	}
}

// Recibe un string 02 y lo vuelve entero
func monthToInt(month string) int {
	switch month {
	case "01":
		return 1
	case "02":
		return 2
	case "03":
		return 3
	case "04":
		return 4
	case "05":
		return 5
	case "06":
		return 6
	case "07":
		return 7
	case "08":
		return 8
	case "09":
		return 9
	case "10":
		return 10
	case "11":
		return 11
	case "12":
		return 12
	default:
		return -1
	}
}

// PeriodLocalToInt receive month in string and return in int string with zeros
// ej Enero -> 01
func PeriodLocalToInt(namePeriod string) (string, error) {
	switch namePeriod {
	case "Enero":
		return "01", nil
	case "Febrero":
		return "02", nil
	case "Marzo":
		return "03", nil
	case "Abril":
		return "04", nil
	case "Mayo":
		return "05", nil
	case "Junio":
		return "06", nil
	case "Julio":
		return "07", nil
	case "Agosto":
		return "08", nil
	case "Septiembre":
		return "09", nil
	case "Octubre":
		return "10", nil
	case "Noviembre":
		return "11", nil
	case "Diciembre":
		return "12", nil
	default:
		return "", fmt.Errorf("no se encontró el periodo")
	}
}
