package test

import (
	"SQLResultExport/service/impl"
	"testing"
)

func TestExportXmlService_Export(t *testing.T) {
	rs := GetTestRS()

	exportService := new(impl.ExportXmlService)
	fileName, err := exportService.Export(rs)
	if err != nil {
		return
	}
	t.Log(fileName)
}
