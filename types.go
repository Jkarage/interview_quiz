package main

import "go.mongodb.org/mongo-driver/mongo"

var cache = make(map[int]*User)

type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
}

type store interface {
	GetUser(id int) (*User, error)
}

type MongoUserStorer struct {
	db    *mongo.Database
	coll  string
	dbHit int
}

type userHandler struct {
	userStore store
}
