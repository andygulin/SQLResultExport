package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportSQLInsertService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportSQLInsertService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}
