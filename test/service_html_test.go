package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportHtmlService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportHtmlService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}
