package impl

import "SQLResultExport/service"

type ExportMarkdownService struct {
}

func (service *ExportMarkdownService) Export(ds service.DataSource) (service.FileName, error) {
	return "", nil
}
