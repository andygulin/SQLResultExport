package impl

import "SQLResultExport/service"

type ExportJsonService struct {
}

func (service *ExportJsonService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
