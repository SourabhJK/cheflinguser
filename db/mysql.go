package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlDetails MySqlDetails
	sqlDbStr     string
	Dbm          *gorp.DbMap
)

func EstablishMysqlSession() {
	sqlDbStr = mysqlDetails.Username + ":" + mysqlDetails.Password + "@tcp(" + mysqlDetails.IP + ":3306)/" + mysqlDetails.DatabaseName
	fmt.Println("mysqlConnctiondetails:", sqlDbStr)
	Db, err := sql.Open("mysql", sqlDbStr)
	if Db == nil || err != nil {
		log.Fatal("Error in establishing db connection")
	}
	Dbm = &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

}

func (mysqlD MySqlDetails) Create(collection string, data interface{}) error {

	db, err := sql.Open("mysql", sqlDbStr)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	query := mysqlD.generateInsertQuery(collection, data)
	_, err = db.Query(query)

	if err != nil {
		fmt.Println("Error in Create:", err)
	}
	return err
}

func (mysqlD MySqlDetails) Read(collection string, findQuery map[string]interface{}) (interface{}, error) {
	var (
		data   []interface{}
		output interface{}
	)
	query := mysqlD.generateSelectQuery(collection, findQuery)
	_, err := Dbm.Select(&data, query)
	if err != nil {
		log.Error("Error in Read:" + err.Error())
	}
	output = data
	return output, err
}

func (mysqlD MySqlDetails) ReadOne(collection string, findQuery map[string]interface{}) (interface{}, error) {
	var data interface{}
	query := mysqlD.generateSelectQuery(collection, findQuery)
	err := Dbm.SelectOne(&data, query)
	if err != nil {
		log.Error("Error in Read:" + err.Error())
	}
	return data, err
}

func (mysqlD MySqlDetails) Update(collection string, selectQuery map[string]interface{}, updateQuery map[string]interface{}) error {
	db, err := sql.Open("mysql", sqlDbStr)
	if err != nil {
		log.Error("Error in setting up the mysql connection:" + err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	query := mysqlD.generateUpdateQuery(collection, selectQuery, updateQuery)

	_, err = db.Query(query)

	if err != nil {
		fmt.Println("Error in Create:", err)
	}

	return err
}

func (mysqlD MySqlDetails) generateSelectQuery(collection string, findQuery map[string]interface{}) string {

	query := "SELECT * FROM " + collection + " WHERE "
	var where []string
	for k, v := range findQuery {
		where = append(where, "`"+k+"` = "+fmt.Sprintf("%v", v))
	}

	query += strings.Join(where, " AND ") + ";"

	return query
}

func (mysqlD MySqlDetails) generateInsertQuery(collection string, data interface{}) string {

	var (
		columns []string
		values  []string
	)
	query := "INSERT INTO " + collection + "  ("
	convertedMap, err := convert(data)

	if err != nil {
		return ""
	}
	for k, v := range convertedMap {
		if k == "Id" {
			continue
		}
		columns = append(columns, k)
		values = append(values, fmt.Sprintf("%v", v))
	}

	query += strings.Join(columns, ", ")
	query += ") VALUES ("
	query += strings.Join(values, ", ")
	query += ");"

	return query
}

func (mysqlD MySqlDetails) generateUpdateQuery(collection string, selectMap map[string]interface{}, updateMap map[string]interface{}) string {

	query := "UPDATE " + collection

	var (
		set   []string
		where []string
	)
	for k, v := range updateMap {
		set = append(set, k+" = "+fmt.Sprintf("%v", v))
	}

	for k, v := range selectMap {
		if k == "Id" && fmt.Sprintf("%v", v) == "0" {
			continue
		}
		where = append(where, "`"+k+"` = "+fmt.Sprintf("%v", v))
	}
	query += " SET " + strings.Join(set, ", ") + "	WHERE " + strings.Join(where, " AND ") + ";"
	fmt.Println("query:", query)
	return query
}

func convert(data interface{}) (map[string]interface{}, error) {
	var output = make(map[string]interface{})

	byteArr, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &output)

	if err != nil {
		return nil, err
	}

	return output, err

}
