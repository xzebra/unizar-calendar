package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

var (
	ErrNoNextVal = errors.New("no existe siguiente valor en la fila")
)

var months = map[string]int{
	"Ene":  1,
	"Feb":  2,
	"Mar":  3,
	"Abr":  4,
	"May":  5,
	"Jun":  6,
	"Jul":  7,
	"Ago":  8,
	"Sept": 9,
	"Oct":  10,
	"Nov":  11,
	"Dic":  12,
}

type Day struct {
	time.Time
	WeekDay string // La, Lb, Ma, Mb...
}

type Semester struct {
	Start string
	Days  []*Day
}

type Schedule map[string][]*ScheduleClass

type ScheduleClass struct {
	Name              string   `json:"nombre"`
	Hours             string   `json:"horas"`
	Partners          []string `json:"compañeros"`
	GoogleMeet        string   `json:"google-meet"`
	PrivateGoogleMeet string   `json:"google-meet-privado"`
}

func getMonth(row []*xlsx.Cell) int {
	month := row[0].Value
	return months[month]
}

func hasNums(row []*xlsx.Cell) bool {
	for _, col := range row {
		for _, c := range col.Value {
			if c >= '0' && c <= '9' {
				return true
			}
		}
	}

	return false
}

func parseStart(rowNum int, sheet *xlsx.Sheet) (year, month int) {
	var err error

	row := sheet.Rows[rowNum].Cells

	year, err = row[0].Int()
	if err != nil {
		panic(err)
	}

	for i := rowNum + 1; i < len(sheet.Rows); i++ {
		row = sheet.Rows[i].Cells
		if row[0].Value != "" {
			month = months[row[0].Value]
			return
		}
	}

	return
}

func nextInt(pos int, cells []*xlsx.Cell) (int, error) {
	for ; pos < len(cells); pos++ {
		if _, err := cells[pos].Int(); err == nil {
			return pos, nil
		}
	}

	return 0, ErrNoNextVal
}

func ensureNextString(pos int, cells []*xlsx.Cell) (int, error) {
	for ; pos < len(cells); pos++ {
		if _, err := cells[pos].Int(); err != nil {
			if cells[pos].Value != "" {
				return pos, nil
			}
		} else {
			return pos, errors.New("entero encontrado en el camino")
		}
	}

	return 0, ErrNoNextVal
}

func parseTable(rowNum int, sheet *xlsx.Sheet) (sem *Semester) {
	year, month := parseStart(rowNum, sheet)
	sem = &Semester{
		Start: fmt.Sprintf("%d/%d", month, year),
	}

	rows := sheet.Rows
	lastDay := 0

	// procesamos las filas de la tabla
	rowNum++
	for rowNum < len(rows) {
		row := rows[rowNum].Cells
		// omitimos el mes, el num de semana y el finde
		row = row[2 : len(row)-2]

		if !hasNums(row) {
			break
		}

		// Iniciamos el procesado de los días. En el momento en el que no cuadre
		// con el formato | numDia | La |, se asumirá que es el final de la
		// tabla y, por lo tanto, del semestre.
		for i := 0; i < len(row); i++ {
			var err error

			i, err = nextInt(i, row)
			if err != nil {
				break
			}

			day, err := row[i].Int()
			if err != nil {
				break
			}

			if day <= lastDay {
				month++
			}
			lastDay = day

			i, err = ensureNextString(i+1, row)
			if err != nil {
				if err == ErrNoNextVal {
					break
				}
				continue
			}

			weekDay := row[i].Value

			date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			sem.Days = append(sem.Days, &Day{
				Time:    date,
				WeekDay: weekDay[:2],
			})
		}

		// avanzamos a la siguiente fila
		rowNum++
	}

	return sem
}

func parseSheets(sheets []*xlsx.Sheet) (sems []*Semester, err error) {
	if len(sheets) == 0 {
		return nil, errors.New("las tablas no contienen información")
	}

	// Cogemos solo la primera hoja, que tiene los horarios del primer
	// y segundo cuatrimestre
	sheet := sheets[0]

	for rowNum, row := range sheet.Rows {
		cells := row.Cells

		// Para identificar las distintas tablas, tenemos que encontrar la fila
		// inicial. Esta siempre empieza por el año, un número, seguidas de
		// "sem".
		if _, err := strconv.Atoi(cells[0].Value); err != nil {
			continue
		}

		if cells[1].Value != "sem" {
			continue
		}

		sems = append(sems, parseTable(rowNum, sheet))
	}

	return
}

type ClassTimes struct {
	Info  *ScheduleClass
	Times []*ClassTime
}

type ClassTime struct {
	Day   time.Time
	Hours string
}

type MergedSchedule map[string]*ClassTimes

// ToOrg exporta a sintaxis de org-mode para ser leido por org-agenda.
//
// Ex:
// * Inteligencia artificial
// <2020-09-29 15:00-16:00 Tue>
// <2020-09-30 15:00-16:00 Tue>
// :STYLE: habit
func (m MergedSchedule) ToOrg() string {
	var out strings.Builder

	i := 0
	newLined := false
	for class, times := range m {
		if len(times.Times) == 0 {
			continue
		}

		if i > 0 && !newLined {
			out.WriteByte('\n')
		}
		i++
		newLined = false

		out.WriteString("* " + class + "\n")
		for _, time := range times.Times {
			out.WriteString(
				fmt.Sprintf("<%s %s>\n",
					time.Day.Format("2006-01-02"),
					time.Hours,
				),
			)
		}
		out.WriteString(":STYLE: habit\n\n")
		newLined = true

		if len(times.Info.Partners) > 0 {
			newLined = false

			out.WriteString("- Compañeros: ")
			out.WriteByte('\n')
			for _, partner := range times.Info.Partners {
				out.WriteString("  + ")
				out.WriteString(partner)
				out.WriteByte('\n')
			}
		}
		if times.Info.GoogleMeet != "" {
			newLined = false

			out.WriteString("- Google Meet: ")
			out.WriteString(times.Info.GoogleMeet)
			out.WriteByte('\n')
		}

		if times.Info.PrivateGoogleMeet != "" {
			newLined = false

			out.WriteString("- Nuestro Google Meet: ")
			out.WriteString(times.Info.PrivateGoogleMeet)
			out.WriteByte('\n')
		}
	}

	return out.String()
}

func mergeSchedule(sch Schedule, sem *Semester) (out MergedSchedule) {
	out = make(MergedSchedule)

	for _, day := range sem.Days {
		schClasses := sch[day.WeekDay]
		for _, schClass := range schClasses {
			// Inicializamos si no lo está
			if _, ok := out[schClass.Name]; !ok {
				out[schClass.Name] = &ClassTimes{
					Info: schClass,
				}
			}

			// Añadimos el horario cuando ya esté inicializado
			out[schClass.Name].Times = append(out[schClass.Name].Times, &ClassTime{
				Day:   day.Time,
				Hours: schClass.Hours,
			})
		}
	}

	return
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	// He usado https://pdftoxls.com/ para convertir el .pdf a .xlsx
	calendar, err := xlsx.OpenFile("./test.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	semesters, err := parseSheets(calendar.Sheets)
	if err != nil {
		log.Fatal(err)
	}

	scheduleFile, err := os.Open("./horario.json")
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(scheduleFile)
	var schedule Schedule
	json.Unmarshal(content, &schedule)

	merged := mergeSchedule(schedule, semesters[0])
	fmt.Println(merged.ToOrg())
}
