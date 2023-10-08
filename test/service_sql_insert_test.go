package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportSQLInsertService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportSQLInsertService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
