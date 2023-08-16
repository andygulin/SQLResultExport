package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportJsonService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportJsonService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}
