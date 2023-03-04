package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserHandler(uStore store) *userHandler {
	return &userHandler{
		userStore: uStore,
	}
}

func (u *userHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
	}

	user, ok := cache[id]
	if !ok {
		// Go to database
		user, err := u.userStore.GetUser(id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		cache[id] = user
		json.NewEncoder(w).Encode(user)
		return

	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err.Error())
	}
	ms := NewMongoStorer(client.Database("asessment"), "users")
	ms.PopulateDB() // Comment this  line for the first trial
	userHandler := NewUserHandler(ms)
	mux.HandleFunc("/", userHandler.HandleGetUser)
	fmt.Println("Started development server at: http://localhost:4096")
	log.Fatal(http.ListenAndServe(":4096", mux))
}
