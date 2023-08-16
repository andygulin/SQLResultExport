package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportCsvService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportCsvService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}