package db

import (
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"
)

var mongoDetails MongoDetails

func GetSession() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{mongoDetails.IP},
		Timeout:  60 * time.Second,
		Database: mongoDetails.DatabaseName,
		Username: mongoDetails.Username,
		Password: mongoDetails.Password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("ERROR: Not Able to Connect to MongoDB:", err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func (mongo MongoDetails) Create(collection string, data interface{}) error {
	s := GetSession()
	c := s.DB("automata").C(collection)
	return c.Insert(data)
}

func (mongo MongoDetails) Read(collection string, findQuery map[string]interface{}) (interface{}, error) {
	var data interface{}
	s := GetSession()
	c := s.DB("automata").C(collection)
	err := c.Find(findQuery).All(&data)
	return data, err
}

func (mongo MongoDetails) ReadOne(collection string, findQuery map[string]interface{}) (interface{}, error) {
	var data interface{}
	s := GetSession()
	c := s.DB("automata").C(collection)
	err := c.Find(findQuery).One(&data)
	return data, err
}

func (mongo MongoDetails) Update(collection string, selectQuery map[string]interface{}, updateQuery map[string]interface{}) error {
	s := GetSession()
	c := s.DB("automata").C(collection)
	return c.Update(selectQuery, bson.M{"$set": updateQuery})
}
