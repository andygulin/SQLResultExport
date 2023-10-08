package test

import . "SQLResultExport/service"

func GetTestRS() ResultSet {
	var rs []map[string]string

	var foo = make(map[string]string)
	foo["name"] = "aaa"
	foo["age"] = "11"
	rs = append(rs, foo)

	var bar = make(map[string]string)
	bar["name"] = "bbb"
	bar["age"] = "22"
	rs = append(rs, bar)

	return rs
}
