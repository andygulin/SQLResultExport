package main

import (
	"SQLResultExport/service"
	"SQLResultExport/service/impl"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"reflect"
)

const (
	ExportTypeExcel     = "excel"
	ExportTypeCSV       = "csv"
	ExportTypeHtml      = "html"
	ExportTypeJson      = "json"
	ExportTypeXml       = "xml"
	ExportTypeSQLInsert = "sql_insert"
	ExportTypeMarkdown  = "md"
)

func init() {
	viper.SetConfigFile("./conf/conf.yaml")
	viper.SetConfigType("yaml")
}

var db *sqlx.DB
var exportType = flag.String("t", ExportTypeExcel, "-t="+ExportTypeExcel)
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
	_ = db.Ping()
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	rows, _ := db.Query(*query)
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

	var exportService service.ExportService
	switch *exportType {
	case ExportTypeExcel:
		exportService = new(impl.ExportExcelService)
		break
	case ExportTypeCSV:
		exportService = new(impl.ExportCsvService)
		break
	case ExportTypeHtml:
		exportService = new(impl.ExportHtmlService)
		break
	case ExportTypeJson:
		exportService = new(impl.ExportJsonService)
		break
	case ExportTypeXml:
		exportService = new(impl.ExportXmlService)
		break
	case ExportTypeSQLInsert:
		exportService = new(impl.ExportSQLInsertService)
		break
	case ExportTypeMarkdown:
		exportService = new(impl.ExportMarkdownService)
		break
	}

	fmt.Println(reflect.TypeOf(exportService))
}
