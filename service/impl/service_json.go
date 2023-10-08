package impl

import (
	. "SQLResultExport/service"
	"SQLResultExport/tool"
	"encoding/json"
	"path/filepath"
)

type ExportJsonService struct {
}

func (service *ExportJsonService) Export(rs ResultSet) (File, error) {
	b, _ := json.MarshalIndent(rs, "", "	")

	var fileName = "Export.json"
	err := tool.WriteFile(fileName, string(b))
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(fileName)
	return File(output), nil
}
