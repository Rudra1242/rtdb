package OrderItem

import (
	mgo "gopkg.in/mgo.v2"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "OrderItem"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

func (m *MongoRepository) GetAllOrderCount() int {
	return 0
}
