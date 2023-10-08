package tool

import (
	. "SQLResultExport/service"
	"container/list"
	"fmt"
	"os"
)

func WriteFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func ToCSVRecords(rs ResultSet) [][]string {
	var records [][]string
	for _, row := range rs {
		var record []string
		for _, value := range row {
			record = append(record, value)
		}
		records = append(records, record)
	}
	return records
}

func ListToArray(list *list.List) []string {
	var str []string
	for e := list.Front(); e != nil; e = e.Next() {
		str = append(str, fmt.Sprint(e.Value))
	}
	return str
}
