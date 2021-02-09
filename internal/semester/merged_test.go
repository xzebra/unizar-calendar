package semester

import (
	"io/ioutil"
	"os"
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

	// Check: Merged map[string][]timeRange
	assert.Equal(t,
		map[string][]timeRange{
			"ssdd": {
				{ // Lx
					Start: time.Date(2021, 1, 11, 17, 0, 0, 0, time.UTC),
					End:   time.Date(2021, 1, 11, 17, 50, 0, 0, time.UTC),
				},
				{ // Lb
					Start: time.Date(2020, 12, 21, 17, 0, 0, 0, time.UTC),
					End:   time.Date(2020, 12, 21, 17, 50, 0, 0, time.UTC),
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
		},
		data.Merged,
	)
}
