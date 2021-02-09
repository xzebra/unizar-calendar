package semester

import (
	"fmt"

	"github.com/xzebra/unizar-calendar/pkg/schedules"
)

type Data struct {
	*Semester
	Schedule schedules.Schedule
	Classes  schedules.ClassNames

	// Merged is an association between class ids and a list of all
	// days when the class should occur.
	Merged map[string][]timeRange
}

func NewData(semester *Semester, files *schedules.SemesterFiles, number int) (*Data, error) {
	parsed, err := schedules.ParseSemesterFiles(files)
	if err != nil {
		return nil, err
	}

	s := &Data{
		Semester: semester,
		Schedule: parsed.Schedule,
		Classes:  parsed.Names,
	}
	s.mergeClassesDays()

	return s, nil
}

func isWildcardDay(dayType string) bool {
	return dayType[1] == 'x'
}

// getSchedGivenDayType returns the schedule associated with a day
// type. If that day type is a non wildcard day, it also returns the
// wildcarded classes (classes that happen all weeks).
func (s *Data) getSchedGivenDayType(dayType string) (sched []*schedules.ScheduleClass) {
	sched = s.Schedule[dayType]
	if !isWildcardDay(dayType) {
		// We have to extract also the wildcard day classes.
		wildcardType := fmt.Sprintf("%c%c", dayType[0], 'x')
		for _, class := range s.Schedule[wildcardType] {
			sched = append(sched, class)
		}
	}
	return
}

func (s *Data) mergeClassesDays() {
	s.Merged = make(map[string][]timeRange)

	// For each type of day
	for day, dayType := range s.Days {
		sched := s.getSchedGivenDayType(dayType)

		// For each class on that day type
		for _, class := range sched {
			// Add times associated to class
			s.Merged[class.ID] = append(s.Merged[class.ID], timeRange{
				Start: class.Start.AddTo(day),
				End:   class.End.AddTo(day),
			})
		}
	}
}
