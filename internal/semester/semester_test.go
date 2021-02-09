package semester_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/gcal"
)

// func NewSemesterFromData(data []byte) (semester *Semester, err error)
func TestNewSemesterFromData(t *testing.T) {
	f, err := os.Open("./testdata/semester.json")
	assert.Nil(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.Nil(t, err)

	sem, err := semester.NewSemesterFromData(data)
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2020, 9, 14, 0, 0, 0, 0, time.UTC), sem.Begin)
	assert.Equal(t, time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC), sem.End)

	assert.Equal(t, gcal.EventDays{
		time.Date(2021, 1, 11, 0, 0, 0, 0, time.UTC):  "Lx",
		time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC):   "Vx",
		time.Date(2020, 12, 22, 0, 0, 0, 0, time.UTC): "Mb",
		time.Date(2020, 12, 21, 0, 0, 0, 0, time.UTC): "Lb",
	}, sem.Days)
}

// ----------------------------------------------------------
// Test initialization
// ----------------------------------------------------------

func TestMain(t *testing.M) {
	// Run everything from project root folder
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("error returning to root folder: ", err)
		os.Exit(1)
	}

	os.Exit(t.Run())
}
