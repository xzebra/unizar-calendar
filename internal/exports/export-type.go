package exports

import (
	"fmt"
)

type ExportType byte

const (
	OrgExport ExportType = iota
	CSVExport
	ICSExport
)

func (e *ExportType) String() string {
	return "export type"
}

func ExportTypes() []string {
	return []string{"org", "csv", "ics"}
}

func (e *ExportType) Set(in string) error {
	switch in {
	case "org":
		*e = OrgExport
	case "csv":
		*e = CSVExport
	case "ics":
		*e = ICSExport
	default:
		return fmt.Errorf("export type `%s` does not exist", in)
	}

	return nil
}
