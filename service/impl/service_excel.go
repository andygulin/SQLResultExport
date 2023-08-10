package impl

import (
	"SQLResultExport/service"
	"github.com/xuri/excelize/v2"
)

type ExportExcelService struct {
}

func (service *ExportExcelService) Export(ds service.DataSource) (service.FileName, error) {
	file := excelize.NewFile()
	sheet, err := file.NewSheet("Sheet")
	if err != nil {
		return "", err
	}

	file.SetActiveSheet(sheet)
	fileName := "Book1.xlsx"
	err = file.SaveAs(fileName)
	if err != nil {
		return "", err
	}
	return "", nil
}
