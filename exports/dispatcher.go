package exports

import (
	"github.com/xzebra/unizar-calendar/semester"
)

func Export(data *semester.Data, exportType ExportType) string {
	switch exportType {
	case OrgExport:
		return toOrgMode(data)
	case GcalExport:
		return toGcal(data)
	}

	return ""
}
