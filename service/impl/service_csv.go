package impl

import (
	. "SQLResultExport/service"
	"SQLResultExport/tool"
	"encoding/csv"
	"os"
	"path/filepath"
)

type ExportCsvService struct {
}

func (service *ExportCsvService) Export(rs ResultSet) (File, error) {
	var fileName = "Export.csv"
	file, err := os.Create(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return "", err
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(tool.ToCSVRecords(rs))
	if err != nil {
		return "", err
	}
	writer.Flush()

	output, _ := filepath.Abs(fileName)
	return File(output), nil
}
