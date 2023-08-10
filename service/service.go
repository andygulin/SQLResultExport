package service

type DataSource []map[string]string
type FileName string
type ExportService interface {
	Export(ds DataSource) (FileName, error)
}
