package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	// "github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "mongodb+srv://baoqdau:dauq123bao@cluster0.yenqw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "chatapp"
const collName = "users"

var ctx context.Context
var collection *mongo.Collection
var client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(database))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Checking connection!")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collName)
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID
	fmt.Println(id)

}

func ReadCollection() {
	fmt.Println("Checking connection!")

	fmt.Println("Still connected")
	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// fmt.Print(cursor)
}

// func WritePi() {
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
// 	id := res.InsertedID
// }
