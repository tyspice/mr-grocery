package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tyspice/mr-grocery/connectionhelper"
	"github.com/tyspice/mr-grocery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getItemsCollection() (*mongo.Collection, gin.HandlerFunc) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error getting mongo client"})
		}
	}

	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ITEMS)
	return collection, nil
}

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		items := []models.GroceryItem{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection, _ := getItemsCollection()

		cur, findError := collection.Find(ctx, bson.D{{}})
		if findError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "problem finding the items collection"})
		}

		for cur.Next(ctx) {
			i := models.GroceryItem{}
			err := cur.Decode(&i)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error decoding cursor"})
			}
			items = append(items, i)
		}

		cur.Close(ctx)
		if len(items) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"status": mongo.ErrNoDocuments})
		}

		c.JSON(http.StatusOK, items)
	}
}

func GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error converting id"})
		}
		var item models.GroceryItem
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection, _ := getItemsCollection()

		if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error decoding cursor"})
		}

		c.JSON(http.StatusOK, item)

	}
}

func CreateOne() {

}

func UpdateOne() {

}

func DeleteOne() {

}
