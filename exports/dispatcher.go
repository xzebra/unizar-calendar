package exports

import (
	"unizar-calendar/semester"
)

func Export(data *semester.Data, exportType ExportType) string {
	switch exportType {
	case OrgExport:
		return toOrgMode(data)
	}

	return ""
}
