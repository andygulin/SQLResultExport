package impl

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExportExcelService struct {
}

func (service *ExportExcelService) Export(ds []map[string]string) (string, error) {
	file := excelize.NewFile()
	sheetName := "Sheet"
	sheet, err := file.NewSheet(sheetName)
	if err != nil {
		return "", err
	}

	cellNum := 1
	for _, row := range ds {
		cellIdx := 65 // 'A'
		for _, value := range row {
			cell := string(rune(cellIdx)) + strconv.Itoa(cellNum)
			_ = file.SetCellValue(sheetName, cell, value)
			cellIdx++
		}
		cellNum++
	}

	file.SetActiveSheet(sheet)
	fileName := "Export.xlsx"
	err = file.SaveAs(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fileName, nil
}
