package impl

import "SQLResultExport/service"

type ExportSQLInsertService struct {
}

func (service *ExportSQLInsertService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
