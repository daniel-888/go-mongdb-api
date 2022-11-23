package main
import (
	"context"
	"log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

var client *mongo.Client
func main () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	MONGO_USERNAME := os.Getenv("MONGO_USERNAME")
	MONGO_PASSWORD := os.Getenv("MONGO_PASSWORD")
	MONGO_HOSTNAME := os.Getenv("MONGO_HOSTNAME")
	MONGO_PORT := os.Getenv("MONGO_PORT")
	MONGO_DB := os.Getenv("MONGO_DB")

	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGO_USERNAME + ":" + MONGO_PASSWORD  + "@" + MONGO_HOSTNAME + ":" + MONGO_PORT + "?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
}