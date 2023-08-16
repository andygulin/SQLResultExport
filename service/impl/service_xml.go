package impl

import (
	"SQLResultExport/tool"
	"bytes"
	"github.com/beevik/etree"
)

type ExportXmlService struct {
}

func (service *ExportXmlService) Export(ds []map[string]string) (string, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	data := doc.CreateElement("data")
	for _, row := range ds {
		rows := data.CreateElement("rows")
		for key, value := range row {
			rows.CreateElement(key).CreateText(value)
		}
	}

	doc.Indent(4)
	buf := new(bytes.Buffer)
	_, _ = doc.WriteTo(buf)
	var fileName = "Export.xml"
	err := tool.WriteFile(fileName, buf.String())
	if err != nil {
		return "", err
	}
	return fileName, nil
}
