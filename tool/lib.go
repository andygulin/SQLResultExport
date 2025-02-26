package tool

import (
	"SQLResultExport/service"
	"container/list"
	"fmt"
	"os"
	"sort"
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

func ToCSVRecords(rs service.ResultSet) [][]string {
	allKeys := make(map[string]struct{})
	for _, m := range rs {
		for key := range m {
			allKeys[key] = struct{}{}
		}
	}
	sortedKeys := make([]string, 0, len(allKeys))
	for key := range allKeys {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	result := make([][]string, 0, len(rs)+1)
	result = append(result, sortedKeys)

	for _, m := range rs {
		row := make([]string, len(sortedKeys))
		for i, key := range sortedKeys {
			row[i] = m[key]
		}
		result = append(result, row)
	}

	return result
}

func ListToArray(list *list.List) []string {
	var str []string
	for e := list.Front(); e != nil; e = e.Next() {
		str = append(str, fmt.Sprint(e.Value))
	}
	return str
}
