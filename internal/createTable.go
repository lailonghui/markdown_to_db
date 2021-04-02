/*
@Time : 2020/12/23 16:17
@Author : lai
@Description :
@File : createTable
*/
package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"markdown_to_db/config"
	"markdown_to_db/db"
	"strings"
)

type TableElement struct {
	Name        string //数据表列名
	Type        string //数据类型
	Description string //数据列描述
	Require     string //是否不为空
	Default     string //默认值
}
type Table struct {
	TableName string //数据表名
	TableNote string //数据表注释
	Body      []*TableElement
}

func Create() {
	var tables []*Table //数据表结构体数据
	tableFile := config.CONF_INSTANCE.FilePath
	b, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Printf("read tableFile err:%v\n", err)
		return
	}
	//加载md文档，生成数据表结构体数组
	tables = initTables(b)

	//生成数据库表
	for _, v := range tables {
		if len(v.Body) == 0 {
			continue
		}
		err := generatedTable(v)
		if err != nil {
			fmt.Printf("db exec failed:%v\n", err)
			break
		}
	}
	//fmt.Println(len(tables))
}

//生成数据库表
func generatedTable(table *Table) error {
	var sql bytes.Buffer
	sql.WriteString(fmt.Sprintf("CREATE TABLE %s\n", table.TableName))
	sql.WriteString(fmt.Sprintf("(\n"))
	var primaryKey bytes.Buffer
	for i, t := range table.Body {
		idString := ""
		isNull := ""
		if i == 0 {
			idString = "NOT NULL DEFAULT id_generator()"
			primaryKey.WriteString(fmt.Sprintf("%s", t.Name))
		}
		if strings.Contains(t.Default, "联合主键") {
			primaryKey.WriteString(fmt.Sprintf(",%s", t.Name))
		}
		if strings.Contains(t.Require, "required") {
			isNull = "NOT NULL"
		}
		sql.WriteString(fmt.Sprintf("%s %s %s %s,\n ", t.Name, t.Type, idString, isNull))
	}
	//fmt.Printf(primaryKey.String())
	sql.WriteString(fmt.Sprintf("CONSTRAINT %s_pkey PRIMARY KEY (%s)\n", table.TableName, primaryKey.String()))

	sql.WriteString(fmt.Sprintf(");\n"))
	sql.WriteString(fmt.Sprintf("COMMENT ON TABLE public.%s IS '%s';\n", table.TableName, table.TableNote))
	for _, t := range table.Body {
		defaultString := ""

		if strings.TrimSpace(t.Default) != "" {
			defaultString = fmt.Sprintf("(%s)", t.Default)
		}
		sql.WriteString(fmt.Sprintf("COMMENT ON COLUMN public.%s.%s IS '%s';\n", table.TableName, t.Name, strings.ReplaceAll(fmt.Sprintf("%s%s", t.Description, defaultString), "*", "")))
	}

	if isExitTable(table.TableName) {
		sql := fmt.Sprintf("DROP TABLE %s", table.TableName)
		err := db.DB.Exec(sql).Error
		if err != nil {
			fmt.Printf("drop table err:%v", err)
		}
	}

	result := db.DB.Exec(sql.String())
	if result.RowsAffected > 0 {
		fmt.Printf("%d row affected,create table success\n", result.RowsAffected)
	}
	return result.Error
}

func isExitTable(tableName string) bool {
	baseSql := "SELECT count(*) FROM information_schema.tables WHERE table_schema =  CURRENT_SCHEMA() AND table_name = ? AND table_type = 'BASE TABLE'"
	var count int
	db.DB.Raw(baseSql, tableName).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}

//加载md文档，生成数据表结构体数组
func initTables(b []byte) []*Table {
	var tables []*Table
	data := string(b)
	dataArr := strings.Split(data, "### ")
	for _, v := range dataArr {
		table := &Table{}
		arr2 := strings.Split(v, "\n")

		//if strings.Contains(arr2[0],"."){
		//	reg,_ := regexp.Compile(`\s+`)
		//	title := reg.Split(arr2[0],-1)
		title := strings.Split(arr2[0], " ")
		table.TableName = title[0]
		table.TableNote = title[1]
		fmt.Println(title[0])
		//fmt.Println(title[1])
		var tableData []string
		for i := 1; i < len(arr2); i++ {
			if strings.Contains(arr2[i], "|") {
				tableData = append(tableData, arr2[i])
			}
		}
		for i := 2; i < len(tableData); i++ {
			arr3 := strings.Split(tableData[i], "|")
			//fmt.Println(arr3)
			body := &TableElement{
				Name:        arr3[1],
				Type:        arr3[2],
				Description: arr3[3],
				Require:     arr3[4],
				Default:     arr3[5],
			}
			table.Body = append(table.Body, body)
		}
		tables = append(tables, table)
	}
	//}
	return tables
}
