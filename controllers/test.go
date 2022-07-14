package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tyspice/mr-grocery/connectionhelper"
	"github.com/tyspice/mr-grocery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTests() gin.HandlerFunc {
	query := bson.D{{}}
	tests := []models.Test{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "no no bad"})
		}
	}

	testCollection := client.Database(connectionhelper.DB).Collection(connectionhelper.TEST)

	cur, findError := testCollection.Find(ctx, query)
	if findError != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "no no bad"})
		}
	}

	for cur.Next(ctx) {
		t := models.Test{}
		err := cur.Decode(&t)
		if err != nil {
			return func(c *gin.Context) {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "no no bad"})
			}
		}
		tests = append(tests, t)
	}

	cur.Close(ctx)
	if len(tests) == 0 {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": mongo.ErrNoDocuments})
		}
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, tests)
	}
}
