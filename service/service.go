package service

import "sort"

type KeyValue struct {
	Key   string
	Value string
}

type ResultSet []map[string]string
type File string

type ExportService interface {
	Export(rs ResultSet) (File, error)
}

func ToKeyValue(rs []map[string]string) [][]KeyValue {
	var sortedResult [][]KeyValue

	for _, row := range rs {
		keys := make([]string, 0, len(row))
		for key := range row {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		var currentPerson []KeyValue
		for _, key := range keys {
			currentPerson = append(currentPerson, KeyValue{
				Key:   key,
				Value: row[key],
			})
		}
		sortedResult = append(sortedResult, currentPerson)
	}
	return sortedResult
}

func SortMaps(rs []map[string]string) []map[string]string {
	var sortedResult []map[string]string
	for _, row := range rs {
		keys := make([]string, 0, len(row))
		for key := range row {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		sortedPerson := make(map[string]string)
		for _, key := range keys {
			sortedPerson[key] = row[key]
		}
		sortedResult = append(sortedResult, sortedPerson)
	}
	return sortedResult
}
