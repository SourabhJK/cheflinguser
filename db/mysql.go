package db

import (
	"strings"
)

var mysqlDetails MySqlDetails

func (mysqlD MySqlDetails) Create(collection string, data interface{}) error{
	s := GetSession()
	c := s.DB("automata").C(collection)
	return c.Insert(data)
}

func (mysqlD MySqlDetails) Read(collection string, findQuery map[string]interface{}) (interface{}, error){
	var data interface{}
	s := GetSession()
	c := s.DB("automata").C(collection)
	err :=  c.Find(findQuery).All(&data)
	return data, err
}

func (mysqlD MySqlDetails) ReadOne(collection string, findQuery map[string]interface{}) (interface{}, error){
	var data interface{}
	s := GetSession()
	c := s.DB("automata").C(collection)
	err :=  c.Find(findQuery).One(&data)
	return data, err
}

func (mysqlD MySqlDetails) Update(collection string, selectQuery map[string]interface{}, updateQuery map[string]interface{}) error{
	s := GetSession()
	c := s.DB("automata").C(collection)
	return c.Update(selectQuery, bson.M{"$set":updateQuery})
}

func (mysqlD MySqlDetails) generateSelectQuery(collection string, findQuery map[string]interface{}) (string){

	query := "SELECT * FROM "+collection+ " WHERE "
	for k,v := range findQuery{
		query += "`"+k + "` = " +fmt.Sprintf("%v", v) + " AND "
	}

	query = query[:len(query)-5]
	query += ";"

return query
}

func (mysqlD MySqlDetails) generateInsertQuery(collection string, data interface{}) (string){

	query := "INSERT INTO "+collection+ "  ("
	convertedMap, err := convert(data)

	if err != nil{
		return ""
	}
	var columns []string
	var values []string
	for k,v := range convertedMap{
		columns = append(columns, k)
		values = append(values, fmt.Sprintf("%v", v))
	}

	query += strings.Join(columns, ", ")
	query += ") VALUES ("
	query += strings.Join(values, ", ")
	query += ");"

	return query
}


func convert(data interface{}) (map[string]interface{}, error){
	var output = make(map[string]interface{})

	byteArr, err := json.Marshal(data)

	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(byteArr, &output)

	if err != nil{
		return nil, err
	}

	return output, err

}
