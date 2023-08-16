package impl

import (
	"SQLResultExport/tool"
	"container/list"
	"fmt"
	"strings"
)

type ExportSQLInsertService struct {
}

func (service *ExportSQLInsertService) Export(ds []map[string]string) (string, error) {
	const tableName = "MY_TABLE"

	var content []string
	for _, row := range ds {
		fields := list.New()
		values := list.New()
		for key, value := range row {
			fields.PushBack("`" + key + "`")
			values.PushBack("'" + value + "'")
		}
		fieldStr := strings.Join(tool.ListToArray(fields), ",")
		valueStr := strings.Join(tool.ListToArray(values), ",")

		sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);\n", tableName, fieldStr, valueStr)
		content = append(content, sql)
	}

	var fileName = "Export.sql"
	err := tool.WriteFile(fileName, strings.Join(content, ""))
	if err != nil {
		return "", err
	}
	return fileName, nil
}
