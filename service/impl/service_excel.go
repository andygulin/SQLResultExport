package impl

import (
	. "SQLResultExport/service"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strconv"
)

type ExportExcelService struct {
}

func (service *ExportExcelService) Export(rs ResultSet) (File, error) {
	rss := ToKeyValue(rs)

	file := excelize.NewFile()
	sheetName := "Sheet"
	sheet, err := file.NewSheet(sheetName)
	if err != nil {
		return "", err
	}

	cellNum := 1
	for _, row := range rss {
		cellIdx := 65 // 'A'
		for _, kv := range row {
			cell := string(rune(cellIdx)) + strconv.Itoa(cellNum)
			_ = file.SetCellValue(sheetName, cell, kv.Value)
			cellIdx++
		}
		cellNum++
	}

	file.SetActiveSheet(sheet)
	fileName := "Export.xlsx"
	err = file.SaveAs(fileName)
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(fileName)
	return File(output), nil
}
