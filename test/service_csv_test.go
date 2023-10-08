package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportCsvService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportCsvService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
