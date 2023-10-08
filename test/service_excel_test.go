package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportExcelService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportExcelService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
