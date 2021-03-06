package semester

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xzebra/unizar-calendar/pkg/schedules"
)

func TestNewData(t *testing.T) {
	f, err := os.Open("./testdata/semester.json")
	assert.Nil(t, err)
	defer f.Close()

	semData, err := ioutil.ReadAll(f)
	assert.Nil(t, err)

	sem, err := NewSemesterFromData(semData)
	assert.Nil(t, err)

	parsed, err := schedules.ParseSemesterFiles(&schedules.SemesterFiles{
		Subjects: "./testdata/asignaturas.csv",
		Schedule: "./testdata/mixed.csv",
	})
	assert.Nil(t, err)

	data, err := NewData(sem, parsed, 1)
	assert.Nil(t, err)

	expected := map[string][]*TimeRange{
		"ssdd": {
			{ // Lb
				Start: time.Date(2020, 12, 21, 17, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 12, 21, 17, 50, 0, 0, time.UTC),
			},
			{ // Lx
				Start: time.Date(2021, 1, 11, 17, 0, 0, 0, time.UTC),
				End:   time.Date(2021, 1, 11, 17, 50, 0, 0, time.UTC),
			},
		},
		"ing_soft": {
			{ // Lb
				Start: time.Date(2020, 12, 21, 18, 10, 0, 0, time.UTC),
				End:   time.Date(2020, 12, 21, 19, 00, 0, 0, time.UTC),
			},
		},
		"ph": {
			{ // Lb
				Start: time.Date(2020, 12, 21, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 12, 21, 14, 0, 0, 0, time.UTC),
			},
		},
		"ia": {
			{ // Mb
				Start: time.Date(2020, 12, 22, 17, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 12, 22, 20, 0, 0, 0, time.UTC),
			},
		},
	}

	// Ignore UUIDs
	for _, days := range data.Merged {
		for _, day := range days {
			day.UUID = ""
		}
	}

	assert.Equal(t, len(expected), len(data.Merged))

	for k, a := range expected {
		b := data.Merged[k]

		sort.Slice(a, func(i, j int) bool {
			return a[i].Start.Before(a[j].Start)
		})

		sort.Slice(b, func(i, j int) bool {
			return b[i].Start.Before(b[j].Start)
		})

		assert.Equal(t, a, b)
	}
}
