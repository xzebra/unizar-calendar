package main

import (
	"errors"
	"fmt"
	"github.com/xzebra/unizar-calendar/internal/exports"
	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/schedules"
	"os"
	"path/filepath"
	"strings"

	"flag"
	"log"
)

var (
	ErrInvalidSemester = errors.New("semester parameter invalid")
)

func cliUsage() {
	fmt.Fprintf(flag.CommandLine.Output(),
		"Usage: %s [options] subjectsFile scheduleFile\n\n",
		filepath.Base(os.Args[0]), // get only filename
	)
	fmt.Printf("Arguments:\n\n")
	fmt.Println("subjectsFile: csv that contains a list of pairs of ID and subject name.")
	fmt.Println(`e.g.:
	class_id;class_name;class_desc
	ph;Proyecto Hardware;""
	ia;Inteligencia Artificial;""
	ing_soft;Ingeniería del Software;"
	This is a description
	It can be multiline
	https://meet.google.com/ijq-umtk-ewp?pli=1&authuser=1"`)

	fmt.Println()

	fmt.Println("scheduleFile: csv that contains the semester schedule.")
	fmt.Println(`e.g.:
	weekday;class_id;start_hour;end_hour
	#SSDD <- This is a comment
	Lx;ssdd;17:00;17:50`)

	fmt.Println()

	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)

	var semesterNum int
	var exportType exports.ExportType = exports.OrgExport

	// Rewrite flag.Usage to show information for positional arguments
	// too.
	flag.Usage = cliUsage

	flag.IntVar(&semesterNum, "s", 1, "semester (1 or 2)")
	flag.Var(&exportType, "e", fmt.Sprintf(
		"export type (%s)",
		strings.Join(exports.ExportTypes(), ","),
	))
	flag.Parse()

	if semesterNum != 1 && semesterNum != 2 {
		log.Fatal(ErrInvalidSemester)
	}

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	subjectsFile := flag.Args()[0]
	scheduleFile := flag.Args()[1]

	if strings.HasPrefix(subjectsFile, "-") ||
		strings.HasPrefix(scheduleFile, "-") {

		flag.Usage()
		os.Exit(1)
	}

	data, err := semester.NewData(
		&schedules.SemesterFiles{
			Subjects: subjectsFile,
			Schedule: scheduleFile,
		}, semesterNum)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(exports.Export(data, exportType))
}