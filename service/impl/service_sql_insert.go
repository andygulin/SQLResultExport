package impl

import (
	. "SQLResultExport/service"
	"SQLResultExport/tool"
	"container/list"
	"fmt"
	"path/filepath"
	"strings"
)

type ExportSQLInsertService struct {
}

func (service *ExportSQLInsertService) Export(rs ResultSet) (File, error) {
	kvs := ToKeyValue(rs)

	const tableName = "MY_TABLE"

	var content []string
	for _, row := range kvs {
		fields := list.New()
		values := list.New()
		for _, kv := range row {
			fields.PushBack("`" + kv.Key + "`")
			values.PushBack("'" + kv.Value + "'")
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

	output, _ := filepath.Abs(fileName)
	return File(output), nil
}
