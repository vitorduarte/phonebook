package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoDBStorage(url string) (ms *MongoStorage, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return
	}

	ctx, _ := context.WithCancel(context.Background())
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	collection := client.Database("phonebook").Collection("contacts")
	ms = &MongoStorage{client: client, collection: collection, ctx: ctx}
	return
}

func (ms *MongoStorage) Create(c Contact) (contactResponse Contact, err error) {
	if c.Name == "" && c.Phone == "" {
		err = fmt.Errorf("Name and phone cannot be empty")
		return
	}

	res, err := ms.collection.InsertOne(ms.ctx, c)
	if err != nil {
		return
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	contactResponse = Contact{Id: id, Name: c.Name, Phone: c.Phone}
	return
}

func (ms *MongoStorage) GetAll() (response []Contact, err error) {
	cur, err := ms.collection.Find(ms.ctx, bson.D{{}})
	if err != nil {
		return
	}

	response, err = buildContactListFromResponse(ms.ctx, cur)
	if err != nil {
		return
	}
	return
}

func buildContactListFromResponse(ctx context.Context, cur *mongo.Cursor) (response []Contact, err error) {
	for cur.Next(ctx) {
		var contact map[string]interface{}
		err = cur.Decode(&contact)
		if err != nil {
			return
		}

		response = append(response, Contact{
			Id:    contact["_id"].(primitive.ObjectID).Hex(),
			Name:  contact["name"].(string),
			Phone: contact["phone"].(string),
		})
	}
	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}

func (ms *MongoStorage) Get(id string) (response Contact, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	findResult := ms.collection.FindOne(ms.ctx, bson.M{"_id": objectId})
	findResult.Decode(&response)
	response.Id = id
	return
}

func (ms *MongoStorage) Update(c Contact) (response Contact, err error) {
	return
}

func (ms *MongoStorage) Delete(id string) (err error) {
	return
}

func (ms *MongoStorage) FindByName(name string) (response []Contact, err error) {
	regexSearch := primitive.Regex{Pattern: name, Options: ""}
	cur, err := ms.collection.Find(ms.ctx, bson.M{"name": regexSearch})
	response, err = buildContactListFromResponse(ms.ctx, cur)
	if err != nil {
		return
	}
	return
}

func (ms *MongoStorage) HealthCheck() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := ms.client.Ping(ctx, nil)
	if err != nil {
		return false
	}
	return true
}
