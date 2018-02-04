package dao

import (
	"log"

	. "main/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "todos"
)

// Establish a connection to database
func (m *TodoDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of todos
func (m *TodoDAO) FindAll() ([]Todo, error) {
	var todos []Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}

// Find a Todo by its id
func (m *TodoDAO) FindById(id string) (Todo, error) {
	var todo Todo
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

// Insert a Todo into database
func (m *TodoDAO) Insert(todo Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}

// Delete an existing Todo
func (m *TodoDAO) Delete(todo Todo) error {
	err := db.C(COLLECTION).Remove(&todo)
	return err
}

// Update an existing Todo
func (m *TodoDAO) Update(todo Todo) error {
	err := db.C(COLLECTION).UpdateId(todo.ID, &todo)
	return err
}
