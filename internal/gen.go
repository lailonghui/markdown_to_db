package internal

import (
	"bytes"
	"db_doc_generator/config"
	"db_doc_generator/db"
	"fmt"
	"io/ioutil"
	"log"
)

type ColumnInfo struct {
	// 列名
	Columnname string
	// 列描述
	Columndesc string
	// 列类型
	Columntype string
}

var ContentBuffer bytes.Buffer = bytes.Buffer{}

func Gen() {
	modules := config.CONF_INSTANCE.Modules

	if len(modules) == 0 {
		log.Fatal("未配置module")
	}
	for _, v := range modules {
		name := v.Name
		tables := v.Tables
		WriteModuleName(name)
		for _, t := range tables {
			WriteTableInfo(t)
		}
	}

	ioutil.WriteFile("table.md", ContentBuffer.Bytes(), 0777)

}

// 输出module名称
func WriteModuleName(moduleName string) {
	ContentBuffer.WriteString(fmt.Sprintf("# %s\n", moduleName))
}

// 输出表格信息
func WriteTableInfo(tableName string) {
	if !isTableExist(tableName) {
		fmt.Println("表:" + tableName + "不存在")
		return
	}
	tableDesc := getTableDesc(tableName)
	ContentBuffer.WriteString(fmt.Sprintf("### %s %s \n", tableName, tableDesc))
	// 输出表格信息
	// 输出表头
	ContentBuffer.WriteString("| Name | Type | Description | Remark | \n")
	ContentBuffer.WriteString("| ---- | ---- | ---- | ---- | \n")
	//输出表格内容
	tableColumns := getTableColumns(tableName)
	for _, tableColumn := range tableColumns {
		ContentBuffer.WriteString(fmt.Sprintf("| %s | %s | %s | %s | \n", tableColumn.Columnname, tableColumn.Columntype, tableColumn.Columndesc, ""))
	}
}

// 判断表是否存在
func isTableExist(table string) bool {
	baseSql := "SELECT count(*) FROM information_schema.tables WHERE table_schema =  CURRENT_SCHEMA() AND table_name = ? AND table_type = 'BASE TABLE'"
	var count int
	db.DB.Raw(baseSql, table).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}

//获取表描述信息
func getTableDesc(table string) string {
	baseSql := `
SELECT CAST
	( obj_description ( relfilenode, 'pg_class' ) AS VARCHAR ) AS COMMENT 
FROM
	pg_class C 
WHERE
	relname = ?`
	var tableDesc string
	db.DB.Raw(baseSql, table).Scan(&tableDesc)
	return tableDesc
}

func getTableColumns(table string) []ColumnInfo {
	baseSql := `
select a.attname as columnName,concat_ws('',t.typname,SUBSTRING(format_type(a.atttypid,a.atttypmod) from '\(.*\)')) as columnType,d.description as columnDesc from pg_class c,pg_attribute a,pg_type t,pg_description d
where c.relname=? and a.attnum>0 and a.attrelid=c.oid and a.atttypid=t.oid and d.objoid=a.attrelid and d.objsubid=a.attnum;
`
	var columnInfos []ColumnInfo
	db.DB.Raw(baseSql, table).Scan(&columnInfos)
	return columnInfos
}
