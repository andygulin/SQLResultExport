package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportJsonService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportJsonService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
