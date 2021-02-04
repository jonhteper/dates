package dates

import (
	"errors"
	"fmt"
	"time"
)

const USLayout = "2006-01-02"

// Obtiene la fecha ─convertida a string─ con layout YYYY-MM-DD (ISO 8601).
func Today() string {
	t := time.Now()
	date, _ := time.Parse(USLayout, t.String()[:10])

	return date.String()[:10]
}

// Obtiene la fecha en formato latino (DD/MM/YYYY). Con el parámetro
// letters se alterna a un layout semejante a "16 de junio de 2020".
func LatinToday(letters bool) string {
	today, _ := LatinDate(Today(), letters)
	return today
}

// CompareDates permite obtener la fecha más antigua, la más lejana o la más próxima en el futuro de un conjunto de
// fechas. Las fechas deben estar contenidas en un slice de strings y el parámetro operation recibe tres valores
// válidos: «minor», «major» y «next».
func CompareDates(operation string, slice []string) (dateResult string, err error) {

	if operation == "minor" || operation == "major" {
		dateResult, err = EvaluateDates(operation, slice)
		if err != nil {
			fmt.Println(err)
			return
		}
		//verify = true
		return
	} else if operation == "next" {
		dateResult, err = GetNextDate(slice)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	err = errors.New("ingrese 'major', 'minor' o 'next' como argumento operation")

	return
}

// EvaluateDates devuelve la fecha más antigua o la más lejana de un grupo de fechas contenidas en un slice de
// strings.
func EvaluateDates(operation string, slice []string) (result string, err error) {
	_ = time.Now
	for i := 0; i <= len(slice)-1; i++ {
		for e := 0; e <= len(slice)-1; e++ {
			if slice[i] != slice[e] {
				if slice[i] != "" && slice[e] != "" {
					dateI, err := time.Parse(USLayout, slice[i])
					if err != nil {
						fmt.Println(err)
						return "", err
					}
					dateE, err := time.Parse(USLayout, slice[e])
					if err != nil {
						fmt.Println(err)
						return "", err
					}
					if operation == "minor" {
						if dateI.Before(dateE) {
							slice[e] = ""
						}
					} else if operation == "major" {
						if dateI.After(dateE) {
							slice[e] = ""
						}
					} else {
						return "", errors.New("ingrese 'major' o 'minor' como argumento operation")
					}
				}
			}
		}
	}
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] != "" {
			return slice[i], err
		}
	}
	return
}

// GetNextDate devuelve la fecha más próxima en el futuro de un conjunto de fechas dispuestas en un slice.
func GetNextDate(slice []string) (result string, err error) {
	today := time.Now()
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] != "" {
			nextDate, err := time.Parse(USLayout, slice[i])
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			if today.After(nextDate) {
				slice[i] = ""
			}
		}
	}
	result, err = EvaluateDates("minor", slice)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return
}

// LatinDate convierte una fecha de formato ISO 8601 (YYYY-MM-DD) al formato DD/MM/YYYY. Con el parámetro
// 'preferablyLetters' se alterna a un layout semejante a "16 de junio de 2020".
func LatinDate(dateISO8601 string, preferablyLetters bool) (dateResult string, err error) {
	year := dateISO8601[:4]
	month := dateISO8601[5:7]
	day := dateISO8601[8:]
	month, err = MonthSpanishName(month)
	if err != nil {
		return
	}

	if preferablyLetters {
		return fmt.Sprintf("%v de %v de %v", day, month, year), nil
	}

	return fmt.Sprintf("%v/%v/%v", day, month, year), nil
}

// MonthSpanishName devuelve el nombre del mes —en español— correspondiente al número del mismo mes.
func MonthSpanishName(number string) (spanishName string, err error) {
	if number == "01" {
		spanishName = "enero"
	} else if number == "02" {
		spanishName = "febrero"
	} else if number == "03" {
		spanishName = "marzo"
	} else if number == "04" {
		spanishName = "abril"
	} else if number == "05" {
		spanishName = "mayo"
	} else if number == "06" {
		spanishName = "junio"
	} else if number == "07" {
		spanishName = "julio"
	} else if number == "08" {
		spanishName = "agosto"
	} else if number == "09" {
		spanishName = "septiembre"
	} else if number == "10" {
		spanishName = "octubre"
	} else if number == "11" {
		spanishName = "noviembre"
	} else if number == "12" {
		spanishName = "diciembre"
	} else {
		err = errors.New("número de mes inválido")
	}
	return
}

// GetLastDayNextMonth devuelve una fecha con formato ISO 8601 con el último día del siguiente mes.
func GetLastDayNextMonth(dateISO8601 string) (newDate string, err error) {
	_ = time.Now()
	dateString := dateISO8601[:8] + "01"
	date, err := time.Parse(USLayout, dateString)
	if err != nil {
		fmt.Println(err)
		return
	}

	date = date.AddDate(0, 2, -1)

	newDate = date.String()
	newDate = newDate[:10]
	return
}

// AddDate emula el comportamiento de la función de la biblioteca estándar con el mismo nombre, con la diferencia que
// dates.AddDate trabaja con strings. Es necesario que el string de la fecha tenga formato ISO 8601
func AddDate(dateISO8601 string, years int, months int, days int) (newDate string, err error) {
	_ = time.Now()
	date, err := time.Parse(USLayout, dateISO8601)
	if err != nil {
		fmt.Println(err)
		return
	}

	date = date.AddDate(years, months, days)

	newDate = date.String()
	newDate = newDate[:10]
	return
}
