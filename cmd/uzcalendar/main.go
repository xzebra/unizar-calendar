package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/xzebra/unizar-calendar/internal/exports"
	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/schedules"

	"flag"
	"log"
)

var (
	ErrInvalidSemester = errors.New("semester parameter invalid")
)

var (
	semesterDataURL = "https://xzebra.github.io/unizar-calendar/data/semester%d.json"
)

func cliUsage() {
	fmt.Fprintf(flag.CommandLine.Output(),
		"Usage: %s [options] subjectsFile scheduleFile\n\n",
		filepath.Base(os.Args[0]), // get only filename
	)
	fmt.Fprintf(flag.CommandLine.Output(), "Arguments:\n\n")
	fmt.Fprintln(flag.CommandLine.Output(), "subjectsFile: csv that contains a list of pairs of ID and subject name.")
	fmt.Fprintln(flag.CommandLine.Output(), `e.g.:
	class_id;class_name;class_desc
	ph;Proyecto Hardware;""
	ia;Inteligencia Artificial;""
	ing_soft;Ingeniería del Software;"
	This is a description
	It can be multiline
	https://meet.google.com/ijq-umtk-ewp?pli=1&authuser=1"`)

	fmt.Fprintln(flag.CommandLine.Output())

	fmt.Fprintln(flag.CommandLine.Output(), "scheduleFile: csv that contains the semester schedule.")
	fmt.Fprintln(flag.CommandLine.Output(), `e.g.:
	weekday;class_id;start_hour;end_hour
	#SSDD <- This is a comment
	Lx;ssdd;17:00;17:50`)

	fmt.Fprintln(flag.CommandLine.Output())

	fmt.Fprintln(flag.CommandLine.Output(), "Options:")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)

	var semesterNum int
	var outputFile string
	var exportType exports.ExportType = exports.OrgExport

	// Rewrite flag.Usage to show information for positional arguments
	// too.
	flag.Usage = cliUsage

	flag.IntVar(&semesterNum, "s", 1, "semester (1 or 2)")
	flag.StringVar(&outputFile, "o", "", "outputFile")
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

	semesterData, err := getSemesterData(semesterNum)
	if err != nil {
		log.Fatal(err)
	}

	sem, err := semester.NewSemesterFromData(semesterData)
	if err != nil {
		log.Fatal(err)
	}

	data, err := semester.NewData(
		sem,
		&schedules.SemesterFiles{
			Subjects: subjectsFile,
			Schedule: scheduleFile,
		}, semesterNum)
	if err != nil {
		log.Fatal(err)
	}

	if outputFile != "" {
		f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("error opening output file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	log.Print(exports.Export(data, exportType))
}

func getSemesterData(semesterNum int) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(semesterDataURL, semesterNum))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
