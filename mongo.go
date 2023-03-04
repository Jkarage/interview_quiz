package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func NewMongoStorer(db *mongo.Database, coll string) *MongoUserStorer {
	return &MongoUserStorer{
		db:   db,
		coll: coll,
	}
}

func (m *MongoUserStorer) PopulateDB() {
	for i := 0; i < 150; i++ {
		m.db.Collection(m.coll).InsertOne(context.TODO(), bson.M{
			"id":       i,
			"username": fmt.Sprintf("user: %v", i),
		})
	}
}

func (m *MongoUserStorer) GetUser(id int) (*User, error) {
	u := &User{}
	res := m.db.Collection(m.coll).FindOne(context.TODO(), bson.M{
		"id": id,
	})
	m.dbHit++
	err := res.Decode(&u)
	return u, err
}
