package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportExcelService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportExcelService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}
