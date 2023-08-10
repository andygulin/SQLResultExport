package impl

import "SQLResultExport/service"

type ExportXmlService struct {
}

func (service *ExportXmlService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
