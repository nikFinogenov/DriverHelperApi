package main

import (
	"DriverHelperApi/mongoDB"
	_ "bytes"
	"context"
	_ "context"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "io"
	"net/http"
	_ "net/http"
	"os"
	_ "strings"
	_ "time"
)

var CLIENT = mongoDB.InitiateMongoClient()

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getCarHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	fmt.Println(mongoDB.GetCar(CLIENT, objId))
}
func setCarHandler(c *gin.Context) {
	var requestBody mongoDB.Car
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.SetCar(CLIENT, requestBody)
}
func updateCarHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var requestBody mongoDB.Car
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.UpdateCar(CLIENT, objId, requestBody)
	//fmt.Println(id)
}
func deleteCarHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	mongoDB.DeleteCar(CLIENT, objId)
}

func getProfileHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	fmt.Println(mongoDB.GetProfile(CLIENT, objId))
}
func setProfileHandler(c *gin.Context) {
	var requestBody mongoDB.Profile
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.SetProfile(CLIENT, requestBody)
}
func updateProfileHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var requestBody mongoDB.Profile
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.UpdateProfile(CLIENT, objId, requestBody)
	//fmt.Println(id)
}
func deleteProfileHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	mongoDB.DeleteProfile(CLIENT, objId)
}

func getEventHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	fmt.Println(mongoDB.GetEvent(CLIENT, objId))
}
func setEventHandler(c *gin.Context) {
	var requestBody mongoDB.Event
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.SetEvent(CLIENT, requestBody)
}
func updateEventHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var requestBody mongoDB.Event
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}
	mongoDB.UpdateEvent(CLIENT, objId, requestBody)
	//fmt.Println(id)
}
func deleteEventHandler(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	mongoDB.DeleteEvent(CLIENT, objId)
}

func testHandler(c *gin.Context) {
	//fmt.Println(mongoDB.GetCars(CLIENT))
	//var cars, _ = json.Marshal(mongoDB.GetCars(CLIENT))
	//var cars = mongoDB.GetCars(CLIENT)
	//var prof = mongoDB.Profile{
	//	FullName: "qwe",
	//	Email:    "qewq",
	//	Phone:    "asdasd",
	//	Avatar:   mongoDB.GetAvatar(CLIENT, "ava.jpg"),
	//}
	//mongoDB.SetProfile(CLIENT, prof)
	collection := CLIENT.Database("data").Collection("profiles")
	doc := collection.FindOne(context.TODO(), bson.M{})
	var profile mongoDB.Profile
	doc.Decode(&profile)
	var Car = mongoDB.Car{
		Make:           "Toyota",
		Model:          "Venza",
		Vin:            "qwe123qwe",
		Year:           2019,
		PurchaseDate:   "31.04.2005",
		Transmission:   "Auto",
		CurrentMileage: 1231234,
		BodyType:       "Crossover",
		ProfileId:      profile.ProfileId,
	}
	mongoDB.SetCar(CLIENT, Car)
	collection = CLIENT.Database("data").Collection("cars")
	doc = collection.FindOne(context.TODO(), bson.M{})
	var car mongoDB.Car
	doc.Decode(&car)
	//WashDoc := mongoDB.Event{
	//	Date:           "24.02.2019",
	//	Additional:     "",
	//	Cost:           35,
	//	CurrentMileage: 234567,
	//	WashStation:    "B-52",
	//	CarId:          &car.CarId,
	//}
	//fmt.Println(WashDoc.CarId)
	//mongoDB.SetEvent(CLIENT, WashDoc)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
func docHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, _ := os.ReadFile("./swaggerui/swagger.json")
	c.Writer.Write(data)
}

func main() {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Static("/swaggerui/", "/DriverHelperApi/swaggerui")

	router.GET("/", homeHandler)
	//router.GET("/swagger", docHandler)

	router.GET("/getCar/:id", getCarHandler)
	router.POST("/addCar", setCarHandler)
	router.PUT("/updateCar/:id", updateCarHandler)
	router.DELETE("/deleteCar/:id", deleteCarHandler)

	router.GET("/getProfile/:id", getProfileHandler)
	router.POST("/addProfile", setProfileHandler)
	router.PUT("/updateProfile/:id", updateProfileHandler)
	router.DELETE("/deleteProfile/:id", deleteProfileHandler)

	router.GET("/getEvent/:id", getEventHandler)
	router.POST("/addEvent", setEventHandler)
	router.PUT("/updateEvent/:id", updateEventHandler)
	router.DELETE("/deleteEvent/:id", deleteEventHandler)
	//router.GET("/test", testHandler)
	//router.GET("/info", tokenTypeValidation(), verifyIdToken(), server.infoHandler)
	//router.GET("/infoUnsafe", server.infoUnsafeHandler)
	//router.POST("/auth", server.authHandler)
	//router.GET("/auth-callback", server.authCallbackGETHandler)
	//router.POST("/auth-callback", server.authCallbackPOSTHandler)
	//router.POST("/auth/provider-link", server.authJsonRedirect)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}
