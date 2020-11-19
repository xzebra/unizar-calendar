package main

import (
	"errors"
	"unizar-calendar/gcal"
	"unizar-calendar/schedules"

	log "github.com/sirupsen/logrus"
)

var (
	ErrNoNextVal = errors.New("no existe siguiente valor en la fila")
)

var (
	daysA      = "eina.unizar.es_hlti3ac2pou7knidr6e6267g4s@group.calendar.google.com"
	daysB      = "eina.unizar.es_ri3mten96cc0s8am0hm080bi94@group.calendar.google.com"
	holidays   = "eina.unizar.es_nvgat6f556c48fmtk7llb5i5l0@group.calendar.google.com"
	exams      = "eina.unizar.es_8g43cd660rntvu09n32g4hsonk@group.calendar.google.com"
	evaluation = "eina.unizar.es_9vuatq1d533o3aoknsej9vbiv8@group.calendar.google.com"
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

// type Semester struct {
// 	Days []Day
// }

// type Day struct {
// 	time.Time
// 	WeekDay string // La, Lb, Ma, Mb...
// }

// type ClassTimes struct {
// 	Info  *schedules.ScheduleClass
// 	Times []*ClassTime
// }

// type ClassTime struct {
// 	Day   time.Time
// 	Hours string
// }

// type MergedSchedule map[string]*ClassTimes

// func NewMergedSchedule(sch schedules.Schedule, names schedules.ClassName, sem Semester) (out MergedSchedule) {
// 	out = make(MergedSchedule)

// 	for _, day := range sem.Days {
// 		schClasses := sch[day.WeekDay]
// 		for _, schClass := range schClasses {
// 			// Inicializamos si no lo está
// 			className := names[schClass.ID]
// 			if _, ok := out[className]; !ok {
// 				out[className] = &ClassTimes{
// 					Info: schClass,
// 				}
// 			}

// 			// Añadimos el horario cuando ya esté inicializado
// 			out[className].Times = append(out[className].Times, &ClassTime{
// 				Day:   day.Time,
// 				Hours: schClass.Hours,
// 			})
// 		}
// 	}

// 	return

// }

// // ToOrg exporta a sintaxis de org-mode para ser leido por org-agenda.
// //
// // Ex:
// // * Inteligencia artificial
// // <2020-09-29 15:00-16:00 Tue>
// // <2020-09-30 15:00-16:00 Tue>
// // :STYLE: habit
// func (m MergedSchedule) ToOrg() string {
// 	var out strings.Builder

// 	i := 0
// 	newLined := false
// 	for class, times := range m {
// 		if len(times.Times) == 0 {
// 			continue
// 		}

// 		if i > 0 && !newLined {
// 			out.WriteByte('\n')
// 		}
// 		i++
// 		newLined = false

// 		out.WriteString("* " + class + "\n")
// 		for _, time := range times.Times {
// 			out.WriteString(
// 				fmt.Sprintf("<%s %s>\n",
// 					time.Day.Format("2006-01-02"),
// 					time.Hours,
// 				),
// 			)
// 		}
// 		out.WriteString(":STYLE: habit\n\n")
// 		newLined = true

// 	}

// 	return out.String()
// }

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	names, err := schedules.ParseClassNames("./classNames.csv")
	if err != nil {
		log.Fatal(err)
	}

	log.Info(names)

	schedule, err := schedules.ParseSchedule("./horario.csv")
	if err != nil {
		log.Fatal(err)
	}

	log.Info(schedule)

	cal, err := gcal.NewGoogleCalendar()
	if err != nil {
		log.Fatal(err)
	}

	types, err := cal.GetCalendarEventDays(daysA)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(types["La"])

	days, err := cal.GetCalendarDays(holidays)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(days)
}
