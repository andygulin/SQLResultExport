package main

import (
	"SQLResultExport/service/impl"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"reflect"
)

func init() {
	viper.SetConfigFile("./conf/conf.yaml")
	viper.SetConfigType("yaml")
}

var db *sqlx.DB
var exportType = flag.String("t", exportTypeExcel, "-t="+exportTypeExcel)
var query = flag.String("q", "", "-q=select * from table")

func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("exportType : %s, query : %s", *exportType, *query))

	if *query == "" {
		fmt.Println("query cannot be empty...")
		return
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbType := viper.GetString("db.type")
	dbHost := viper.GetString("db.host")
	dbName := viper.GetString("db.name")
	dbPort := viper.GetInt("db.port")
	dbUser := viper.GetString("db.user")
	dbPass := viper.GetString("db.pass")

	ds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ = sqlx.Connect(dbType, ds)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	rows, err := db.Query(*query)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}

	res := make([]map[string]string, 0)
	for rows.Next() {
		_ = rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}
		res = append(res, row)
	}

	fileName := invokeService(*exportType, res)
	fmt.Printf("Export File : %s\n", fileName)
}

const (
	exportTypeExcel     = "excel"
	exportTypeCSV       = "csv"
	exportTypeHtml      = "html"
	exportTypeJson      = "json"
	exportTypeXml       = "xml"
	exportTypeSQLInsert = "sql_insert"
)

func invokeService(exportType string, res []map[string]string) string {
	serviceMap := make(map[string]reflect.Type, 7)
	serviceMap[exportTypeExcel] = reflect.TypeOf(impl.ExportExcelService{})
	serviceMap[exportTypeCSV] = reflect.TypeOf(impl.ExportCsvService{})
	serviceMap[exportTypeHtml] = reflect.TypeOf(impl.ExportHtmlService{})
	serviceMap[exportTypeJson] = reflect.TypeOf(impl.ExportJsonService{})
	serviceMap[exportTypeXml] = reflect.TypeOf(impl.ExportXmlService{})
	serviceMap[exportTypeSQLInsert] = reflect.TypeOf(impl.ExportSQLInsertService{})

	var t = serviceMap[exportType]
	obj := reflect.New(t)
	method := obj.MethodByName("Export")
	args := []reflect.Value{reflect.ValueOf(res)}
	rets := method.Call(args)
	fileName := rets[0].String()
	tt := rets[1].Interface()
	if tt != nil {
		panic(tt.(error))
	}
	return fileName
}
