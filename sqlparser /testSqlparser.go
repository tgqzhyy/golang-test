package main

import (
	"encoding/json"
	"fmt"
	"github.com/ha666/golibs"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
)

/**
use this library github.com/xwb1989/sqlparser

*/
func main() {
	b, err := ioutil.ReadFile("abc.sql")
	if err != nil {
		fmt.Println("Read File ERR::", err.Error())
		return
	}
	//fmt.Println(string(b))
	_createTable := golibs.SliceByteToString(b)
	//fmt.Println(_createTable)
	{

		stmt, err := sqlparser.Parse(_createTable)
		if err != nil {
			fmt.Println("Parser Sql ERR:", err.Error())
		}
		//fmt.Println(golibs.ToJson(stmt))

		var _sqlStructure createTableSqlStructure
		err = json.Unmarshal(golibs.StringToSliceByte(golibs.ToJson(stmt)), &_sqlStructure)
		if err != nil {
			fmt.Println("Parser Sql ERR:", err.Error())
		}
		fmt.Println(golibs.FormatJson(_sqlStructure))
	}
}

type createTableSqlStructure struct {
	Action string `json:"Action"`
	Table  struct {
		Name      string `json:"Name"`
		Qualifier string `json:"Qualifier"`
	} `json:"Table"`
	NewName struct {
		Name      string `json:"Name"`
		Qualifier string `json:"Qualifier"`
	} `json:"NewName"`
	IfExists  bool `json:"IfExists"`
	TableSpec struct {
		Columns []struct {
			Name string `json:"Name"`
			Type struct {
				Type          string      `json:"Type"`
				NotNull       bool        `json:"NotNull"`
				Autoincrement bool        `json:"Autoincrement"`
				Default       interface{} `json:"Default"`
				OnUpdate      interface{} `json:"OnUpdate"`
				Comment       struct {
					Type int    `json:"Type"`
					Val  string `json:"Val"`
				} `json:"Comment"`
				Length struct {
					Type int    `json:"Type"`
					Val  string `json:"Val"`
				} `json:"Length"`
				Unsigned   bool        `json:"Unsigned"`
				Zerofill   bool        `json:"Zerofill"`
				Scale      interface{} `json:"Scale"`
				Charset    string      `json:"Charset"`
				Collate    string      `json:"Collate"`
				EnumValues interface{} `json:"EnumValues"`
				KeyOpt     int         `json:"KeyOpt"`
			} `json:"Type"`
		} `json:"Columns"`
		Indexes []struct {
			Info struct {
				Type    string `json:"Type"`
				Name    string `json:"Name"`
				Primary bool   `json:"Primary"`
				Spatial bool   `json:"Spatial"`
				Unique  bool   `json:"Unique"`
			} `json:"Info"`
			Columns []struct {
				Column string      `json:"Column"`
				Length interface{} `json:"Length"`
			} `json:"Columns"`
			Options interface{} `json:"Options"`
		} `json:"Indexes"`
		Options string `json:"Options"`
	} `json:"TableSpec"`
	PartitionSpec interface{} `json:"PartitionSpec"`
	VindexSpec    interface{} `json:"VindexSpec"`
	VindexCols    interface{} `json:"VindexCols"`
}
