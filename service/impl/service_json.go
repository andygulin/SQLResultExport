package impl

import (
	"SQLResultExport/tool"
	"encoding/json"
)

type ExportJsonService struct {
}

func (service *ExportJsonService) Export(ds []map[string]string) (string, error) {
	b, _ := json.MarshalIndent(ds, "", "	")

	var fileName = "Export.json"
	err := tool.WriteFile(fileName, string(b))
	if err != nil {
		return "", err
	}
	return fileName, nil
}
