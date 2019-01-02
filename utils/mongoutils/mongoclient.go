package mongoutils

import (
	"rtdb/utils/loggerutils"

	mgo "gopkg.in/mgo.v2"
)

type MongoAuthObject struct {
	DBname   string
	Username string
	Password string
}

//MongoSession stores mongo session
var MongoSession *mgo.Session

//RegisterMongoSession creates a new mongo session
func RegisterMongoSession(mongoURI string, connectionType string, mongoauth *MongoAuthObject) (*mgo.Session, error) {
	logger := loggerutils.GetLogger()
	var err error
	MongoSession, err = mgo.Dial(mongoURI)
	if err != nil {
		logger.Errorf("Error in Connecting Mongo")
		panic(err)
	}
	if connectionType == "authenticated" {
		var err2 error
		err2 = MongoSession.DB(mongoauth.DBname).Login(mongoauth.Username, mongoauth.Password)
		if err2 != nil {
			logger.Errorf("Not able to login to Mongo with username and password")
			panic(err2)
		}
	}
	return MongoSession, nil
}
