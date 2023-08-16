package test

func GetTestDS() []map[string]string {
	var ds []map[string]string

	var foo = make(map[string]string)
	foo["name"] = "aaa"
	foo["age"] = "11"
	ds = append(ds, foo)

	var bar = make(map[string]string)
	bar["name"] = "bbb"
	bar["age"] = "22"
	ds = append(ds, bar)

	return ds
}
