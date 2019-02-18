package main

import (
	"log"

	"github.com/globalsign/mgo"
)

func main() {
	var err error
	//mongoURL := os.Getenv("MONGO_URL")
	mongoURL := "test_user:test_secret@127.0.0.1:27017"
	if mongoURL == "" {
		log.Fatal("MONGO_URL not provided")
	}
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	err = session.DB("").AddUser("test_user", "test_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("followings").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
