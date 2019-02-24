package storage

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/zhongweili/user-timeline/pkg/core"
)

const (
	collectionName = "followings"
)

func GetCollectionName() string {
	return collectionName
}

type MongoRepository struct {
	logger  *log.Logger
	session *mgo.Session
}

// Find fetches a following from mongo according to the query criteria provided.
func (r MongoRepository) Find(uId string) (*core.Following, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var following core.Following
	err := coll.Find(bson.M{"userId": following.UserID, "uId": uId}).One(&following)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return &following, nil
}

// FindAll fetches all followings from the database. YES.. ALL! be careful.
func (r MongoRepository) FindAll(selector map[string]interface{}) ([]*core.Following, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var followings []*core.Following
	err := coll.Find(selector).All(&followings)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return followings, nil
}

// Delete deletes a following from mongo according to the query criteria provided.
func (r MongoRepository) Delete(following *core.Following) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Remove(bson.M{"userId": following.UserID, "uId": following.UID})
}

// Update updates an following.
func (r MongoRepository) Update(following *core.Following) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Update(bson.M{"userId": following.UserID, "uId": following.UID}, following)
}

// Create followings in the database.
func (r MongoRepository) Create(followings ...*core.Following) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	for _, following := range followings {
		_, err := coll.Upsert(bson.M{"userId": following.UserID, "uId": following.UID}, following)
		if err != nil {
			return err
		}
	}

	return nil
}

// Count counts documents for a given collection
func (r MongoRepository) Count() (int, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)
	return coll.Count()
}

// NewMongoSession dials mongodb and creates a session.
func newMongoSession() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL not provided")
	}
	return mgo.Dial(mongoURL)
}

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB] ", 0)
}

func NewMongoRepository() core.Repository {
	logger := newMongoRepositoryLogger()
	session, err := newMongoSession()
	if err != nil {
		logger.Fatalf("Could not connect to the database: %v\n", err)
	}

	return MongoRepository{
		session: session,
		logger:  logger,
	}
}
