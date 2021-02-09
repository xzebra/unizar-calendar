package schedules

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------
// Test functions
// ----------------------------------------------------------

func TestParseClassNames(t *testing.T) {
	expected := ClassNames{
		"ph": &ClassName{ID: "ph", Name: "Proyecto Hardware", Desc: ""},
		"ia": &ClassName{ID: "ia", Name: "Inteligencia Artificial", Desc: ""},
		"ing_soft": &ClassName{ID: "ing_soft", Name: "Ingeniería del Software",
			Desc: "\nProblemas: https://meet.google.com/ijq-umtk-ewp?pli=1&authuser=1",
		},
		"sist_info": &ClassName{ID: "sist_info", Name: "Sistemas de Información", Desc: ""},
		"ssdd":      &ClassName{ID: "ssdd", Name: "Sistemas Distribuidos", Desc: ""},
	}

	f, err := os.Open("./testdata/asignaturas.csv")
	assert.Nil(t, err)
	defer f.Close()

	returned, err := ParseClassNames(f)
	assert.Nil(t, err)

	assert.Equal(t, expected, returned)
}

func TestParseSchedule_Classes(t *testing.T) {
	class1 := &ScheduleClass{
		Weekday: "Lx",
		ID:      "ssdd",
		Start:   hour{17, 0},
		End:     hour{17, 50},
	}
	class2 := &ScheduleClass{
		Weekday: "Lb",
		ID:      "ing_soft",
		Start:   hour{18, 10},
		End:     hour{19, 00},
	}
	expected := Schedule{
		"Lx": {class1},
		"Lb": {class2},
	}

	f2, err := os.Open("./testdata/problemas.csv")
	assert.Nil(t, err)
	defer f2.Close()

	returned, err := ParseSchedule(f2)
	assert.Nil(t, err)

	assert.Equal(t, expected, returned)
}

func TestParseSchedule_Practical(t *testing.T) {
	class1 := &ScheduleClass{
		Weekday:     "Xx",
		ID:          "ph",
		Start:       hour{10, 0},
		End:         hour{14, 0},
		IsPractical: true,
	}
	class2 := &ScheduleClass{
		Weekday:     "Ja",
		ID:          "ia",
		Start:       hour{17, 0},
		End:         hour{20, 0},
		IsPractical: true,
	}
	expected := Schedule{
		"Xx": {class1},
		"Ja": {class2},
	}

	f2, err := os.Open("./testdata/practicas.csv")
	assert.Nil(t, err)
	defer f2.Close()

	returned, err := ParseSchedule(f2)
	assert.Nil(t, err)

	assert.Equal(t, expected, returned)
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
