package impl

import (
	"SQLResultExport/tool"
	"encoding/csv"
	"os"
)

type ExportCsvService struct {
}

func (service *ExportCsvService) Export(ds []map[string]string) (string, error) {
	var fileName = "Export.csv"
	file, err := os.Create(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return "", err
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(tool.ToCSVRecords(ds))
	if err != nil {
		return "", err
	}
	writer.Flush()

	return fileName, nil
}
