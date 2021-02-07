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

	returned, err := ParseClassNames("./testdata/asignaturas.csv")
	assert.Nil(t, err)

	assert.Equal(t, expected, returned)
}

func TestParseSchedule(t *testing.T) {
	class1 := &ScheduleClass{
		Weekday: "Lx",
		ID:      "ssdd",
		Start:   hour{17, 0},
		End:     hour{17, 50},
	}
	class2 := &ScheduleClass{
		Weekday: "Lx",
		ID:      "ing_soft",
		Start:   hour{18, 10},
		End:     hour{19, 00},
	}
	expected := Schedule{
		"La": {class1, class2},
		"Lb": {class1, class2},
	}

	returned, err := ParseSchedule("./testdata/problemas.csv")
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
