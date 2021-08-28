package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/vitorduarte/phonebook/internal/contact"
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

func NewMongoStorage(url string) (ms *MongoStorage, err error) {
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

func (ms *MongoStorage) Create(c contact.Contact) (contactResponse contact.Contact, err error) {
	if c.Name == "" && c.Phone == "" {
		err = fmt.Errorf("Name and phone cannot be empty")
		return
	}

	res, err := ms.collection.InsertOne(ms.ctx, c)
	if err != nil {
		return
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	contactResponse = contact.Contact{Id: id, Name: c.Name, Phone: c.Phone}
	return
}

func (ms *MongoStorage) GetAll() (response []contact.Contact, err error) {
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

func buildContactListFromResponse(ctx context.Context, cur *mongo.Cursor) (response []contact.Contact, err error) {
	for cur.Next(ctx) {
		var curContact map[string]interface{}
		err = cur.Decode(&curContact)
		if err != nil {
			return
		}

		response = append(response, contact.Contact{
			Id:    curContact["_id"].(primitive.ObjectID).Hex(),
			Name:  curContact["name"].(string),
			Phone: curContact["phone"].(string),
		})
	}
	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}

func (ms *MongoStorage) Get(id string) (contact contact.Contact, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	findResult := ms.collection.FindOne(ms.ctx, bson.M{"_id": objectId})
	if findResult.Err() != nil {
		err = fmt.Errorf("contact with id: %s does not exist on database", id)
		return
	}

	findResult.Decode(&contact)
	contact.Id = id
	return
}

func (ms *MongoStorage) Update(c contact.Contact) (contactResponse contact.Contact, err error) {
	id, err := primitive.ObjectIDFromHex(c.Id)
	if err != nil {
		return
	}

	updateResponse, err := ms.collection.UpdateOne(
		ms.ctx,
		bson.M{"_id": id},
		bson.D{{"$set", c}},
	)
	if err != nil {
		return
	}
	if updateResponse.ModifiedCount == 0 {
		err = fmt.Errorf("contact with id: %s does not exist on database", c.Id)
	}

	contactResponse = c
	return
}

func (ms *MongoStorage) Delete(id string) (err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	deleteResponse, err := ms.collection.DeleteOne(ms.ctx, bson.M{"_id": objectId})
	if err != nil {
		return
	}
	if deleteResponse.DeletedCount == 0 {
		err = fmt.Errorf("contact with id: %s does not exist on database", id)
	}

	return
}

func (ms *MongoStorage) FindByName(name string) (response []contact.Contact, err error) {
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
