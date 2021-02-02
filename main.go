package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unizar-calendar/exports"
	"unizar-calendar/schedules"
	"unizar-calendar/semester"

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
	fmt.Println("Arguments:\n")
	fmt.Println("subjectsFile: csv that contains a list of pairs of ID and subject name.")
	fmt.Println(`e.g.:
	class_id;class_name
	ph;Proyecto Hardware
	ia;Inteligencia Artificial
	ing_soft;Ingeniería del Software
`)

	fmt.Println("scheduleFile: csv that contains the semester schedule.")
	fmt.Println(`e.g.:
	weekday;class_id;start_hour;end_hour;desc
	#SSDD <- This is a comment
	Lx;ssdd;17:00;17:50;"
	This is a description
	It can be multiline
	https://meet.google.com/ijq-umtk-ewp?pli=1&authuser=1"
`)

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
	flag.Var(&exportType, "e", "export type (text, org)")
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
