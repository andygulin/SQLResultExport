package service

type ExportService interface {
	Export(ds []map[string]string) (string, error)
}
