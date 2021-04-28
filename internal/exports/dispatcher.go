package exports

import (
	"github.com/xzebra/unizar-calendar/internal/semester"
)

func Export(data *semester.Data, exportType ExportType) string {
	switch exportType {
	case OrgExport:
		return toOrgMode(data)
	case CSVExport:
		return toGcalCSV(data)
	case ICSExport:
		return toGcalICS(data)
	}

	return ""
}
