package exports

import (
	"fmt"
)

type ExportType byte

const (
	OrgExport ExportType = iota
)

func (e *ExportType) String() string {
	return "export type"
}

func (e *ExportType) Set(in string) error {
	switch in {
	case "org":
		*e = OrgExport
	}

	return fmt.Errorf("export type `%s` does not exist", in)
}
