# 根据SQL结果集导出不同格式的数据

### 暂时只支持MySQL数据库

## Code

- [Excel](service/impl/service_excel.go)
- [CSV](service/impl/service_csv.go)
- [Json](service/impl/service_json.go)
- [Xml](service/impl/service_xml.go)
- [SQL INSERT](service/impl/service_sql_insert.go)
- [Html](service/impl/service_html.go)

## Test

- [Excel](test/service_excel_test.go)
- [CSV](test/service_csv_test.go)
- [Json](test/service_json_test.go)
- [Xml](test/service_xml_test.go)
- [SQL INSERT](test/service_sql_insert_test.go)
- [Html](test/service_html_test.go)

## Usage

- [DataSource](conf/conf.yaml)

```shell
go build
```

```shell
Excel
./SQLResultExport -t=excel -q="select * from table"

CSV
./SQLResultExport -t=csv -q="select * from table"

Json
./SQLResultExport -t=json -q="select * from table"

Xml
./SQLResultExport -t=xml -q="select * from table"

SQL INSERT
./SQLResultExport -t=sql_insert -q="select * from table"

Html
./SQLResultExport -t=html -q="select * from table"
```
