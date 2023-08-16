package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportXmlService_Export(t *testing.T) {
	ds := GetTestDS()

	exportService := new(impl.ExportXmlService)
	fileName, err := exportService.Export(ds)
	if err != nil {
		return
	}
	t.Log(fileName)
}
