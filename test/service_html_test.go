package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportHtmlService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportHtmlService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
