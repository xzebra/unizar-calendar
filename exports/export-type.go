package exports

import (
	"fmt"
)

type ExportType byte

const (
	OrgExport ExportType = iota
	GcalExport
)

func (e *ExportType) String() string {
	return "export type"
}

func ExportTypes() []string {
	return []string{"org", "gcal"}
}

func (e *ExportType) Set(in string) error {
	switch in {
	case "org":
		*e = OrgExport
	case "gcal":
		*e = GcalExport
	default:
		return fmt.Errorf("export type `%s` does not exist", in)
	}

	return nil
}
