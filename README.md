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

```sql
create table `tb`
(
    id         int auto_increment,
    name       varchar(20) default ''    null,
    age        int         default 0     null,
    created_at datetime    default NOW() null,
    sex        enum ('男','女','未知') null,
	constraint table_pk_id
		primary key (id)
);

INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (1, '小明', 11, '2023-10-08 11:31:42', '男');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (2, '小红', 12, '2023-10-08 11:31:56', '女');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (3, '小强', 13, '2023-10-08 11:32:08', '男');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (4, '小米', 14, '2023-10-08 11:32:20', '未知');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (5, '小乔', 15, '2023-10-08 11:32:39', '女');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (6, '小车', 16, '2023-10-08 11:32:40', '女');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (7, '小儿', 17, '2023-10-08 11:33:21', '男');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (8, '小宝', 18, '2023-10-08 11:33:22', '女');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (9, '小白', 19, '2023-10-08 11:33:23', '男');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (10, '小茹', 20, '2023-10-08 11:33:24', '女');
INSERT INTO `tb` (id, name, age, created_at, sex) VALUES (11, '小欧', 21, '2023-10-08 11:33:38', '未知');
```

```shell
go build
```

```shell
Excel
./SQLResultExport -t=excel -q="select * from tb"

CSV
./SQLResultExport -t=csv -q="select * from tb"

Json
./SQLResultExport -t=json -q="select * from tb"

Xml
./SQLResultExport -t=xml -q="select * from tb"

SQL INSERT
./SQLResultExport -t=sql_insert -q="select * from tb"

Html
./SQLResultExport -t=html -q="select * from tb"
```
