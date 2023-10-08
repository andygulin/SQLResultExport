package service

type ResultSet []map[string]string
type File string

type ExportService interface {
	Export(rs ResultSet) (File, error)
}
