package handlers

import (
	"context"
	"crud-mongo-gcp/pkg/models"
	"encoding/json"
	"net/http"
	"time"

	"crud-mongo-gcp/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	client, err := database.ConnectDB()
	if err != nil {
		http.Error(w, `{"message":"Failed to connect database"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	collection := client.Database("GCP").Collection("items")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		http.Error(w, `{"message":"Failed to insert item"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	client, err := database.ConnectDB()
	if err != nil {
		http.Error(w, `{"message":"Failed to connect database"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("GCP").Collection("items")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, `{"message":"Failed to get items"}`, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var items []models.Item
	for cursor.Next(ctx) {
		var item models.Item
		cursor.Decode(&item)
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}
