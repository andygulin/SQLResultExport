package impl

import "SQLResultExport/service"

type ExportCsvService struct {
}

func (service *ExportCsvService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
