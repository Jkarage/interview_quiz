package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_HandleGetUser(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Connection to Database failed, Test failed")
	}
	mongoStore := NewMongoStorer(client.Database("asessment"), "users")
	mongoStore.PopulateDB()
	userHandler := NewUserHandler(mongoStore)
	ts := httptest.NewServer(http.HandlerFunc(userHandler.HandleGetUser))

	nreq := 1001
	for i := 1; i < nreq; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
		resp, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}
		user := &User{}
		err = json.NewDecoder(resp.Body).Decode(user)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("Request Number: %d %+v\n", i, user)
	}
	fmt.Println("The number of requests we had: ", nreq)
	fmt.Println("The number of times requests hit db: ", mongoStore.dbHit)
}
