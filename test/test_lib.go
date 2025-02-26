package test

import (
	. "SQLResultExport/service"
	"math/rand"
	"strconv"
	"time"
)

func GetTestRS() ResultSet {
	var rs []map[string]string

	for i := 0; i < 10; i++ {
		var foo = make(map[string]string)
		foo["name"] = randomString(10)
		foo["age"] = strconv.Itoa(randomNumber())
		rs = append(rs, foo)
	}

	return rs
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

func randomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return rng.Intn(90) + 10
}
