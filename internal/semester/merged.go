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

func NewData(semester *Semester, parsed *schedules.ParsedSemesterFiles, number int) (*Data, error) {
	s := &Data{
		Semester: semester,
		Schedule: parsed.Schedule,
		Classes:  parsed.Names,
	}
	return s, s.mergeClassesDays()
}

func isWildcardDay(dayType string) bool {
	return dayType[1] == 'x'
}

func removePracticalClasses(sched []*schedules.ScheduleClass) []*schedules.ScheduleClass {
	var out []*schedules.ScheduleClass

	for _, class := range sched {
		if !class.IsPractical {
			out = append(out, class)
		}
	}

	return out
}

// getSchedGivenDayType returns the schedule associated with a day
// type. If that day type is a non wildcard day, it also returns the
// wildcarded classes (classes that happen all weeks).
func (s *Data) getSchedGivenDayType(dayType string) (sched []*schedules.ScheduleClass) {
	sched = s.Schedule[dayType]
	if isWildcardDay(dayType) { // Day with classes but not practical classes
		sched = removePracticalClasses(sched)
	} else { // Day with practical classes
		// We have to extract also the wildcard day classes.
		wildcardType := fmt.Sprintf("%c%c", dayType[0], 'x')
		for _, class := range s.Schedule[wildcardType] {
			sched = append(sched, class)
		}
	}
	return
}

func (s *Data) mergeClassesDays() error {
	s.Merged = make(map[string][]timeRange)

	// For each type of day
	for day, dayType := range s.Days {
		sched := s.getSchedGivenDayType(dayType)

		// For each class on that day type
		for _, class := range sched {
			// Check if class id exists
			if _, exists := s.Classes[class.ID]; !exists {
				return fmt.Errorf("used non existing class id `%s`", class.ID)
			}

			// Add times associated to class
			s.Merged[class.ID] = append(s.Merged[class.ID], timeRange{
				UUID:  class.UUID,
				Start: class.Start.AddTo(day),
				End:   class.End.AddTo(day),
			})
		}
	}

	return nil
}
