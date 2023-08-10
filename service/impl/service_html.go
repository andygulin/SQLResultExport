package impl

import "SQLResultExport/service"

type ExportHtmlService struct {
}

func (service *ExportHtmlService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
