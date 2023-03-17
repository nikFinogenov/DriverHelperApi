// fetch('http://localhost:4112/auth', {method: 'POST', body: JSON.stringify({provider: 'google'})})
package mongoDB

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Car struct {
	Make           string              `bson:"make,omitempty" json:"make,omitempty"`
	Model          string              `bson:"model,omitempty" json:"model,omitempty"`
	Vin            string              `bson:"vin,omitempty" json:"vin,omitempty"`
	Year           uint16              `bson:"year,omitempty" json:"year,omitempty"`
	PurchaseDate   string              `bson:"purchase_date,omitempty" json:"purchaseDate,omitempty"`
	Transmission   string              `bson:"transmission,omitempty" json:"transmission,omitempty"`
	CurrentMileage uint32              `bson:"current_mileage,omitempty" json:"currentMileage,omitempty"`
	BodyType       string              `bson:"body_type,omitempty" json:"bodyType,omitempty"`
	CarId          *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ProfileId      *primitive.ObjectID `bson:"p_id,omitempty" json:"p_id,omitempty"`
}
type Profile struct {
	FullName  string              `bson:"full_name,omitempty" json:"fullName,omitempty"`
	Email     string              `bson:"email,omitempty" json:"email,omitempty"`
	Avatar    string              `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Phone     string              `bson:"phone,omitempty" json:"phone,omitempty"`
	ProfileId *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}
type Event struct {
	Date           string              `bson:"date,omitempty" json:"date,omitempty"`
	CurrentMileage uint32              `bson:"current_mileage,omitempty" json:"currentMileage,omitempty"`
	Additional     string              `bson:"additional,omitempty" json:"additional,omitempty"`
	Cost           uint16              `bson:"cost,omitempty" json:"cost,omitempty"`
	Tags           []string            `bson:"tags,omitempty" json:"tags,omitempty"`
	EventType      string              `bson:"event_type,omitempty" json:"eventType,omitempty"`
	PricePer1L     uint16              `bson:"price_per_1_l,omitempty" json:"pricePer1L,omitempty"`
	GasStation     string              `bson:"gas_station,omitempty" json:"gasStation,omitempty"`
	FuelType       string              `bson:"fuel_type,omitempty" json:"fuelType,omitempty"`
	FuelAmount     uint16              `bson:"fuel_amount,omitempty" json:"fuelAmount,omitempty"`
	WashStation    string              `bson:"wash_station,omitempty" json:"washStation,omitempty"`
	RepairStation  string              `bson:"repair_station,omitempty" json:"repairStation,omitempty"`
	PartsList      []string            `bson:"parts_list,omitempty" json:"partsList,omitempty"`
	CarId          *primitive.ObjectID `bson:"c_id,omitempty" json:"c_id,omitempty"`
	EventId        *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

//	func GetCars(cl *mongo.Client) interface{} {
//		filter := bson.D{}
//		coll := cl.Database("data").Collection("cars")
//		cursor, err := coll.Find(context.TODO(), filter)
//		if err != nil {
//			log.Fatal(err)
//		}
//		var cars []Car
//		if err = cursor.All(context.TODO(), &cars); err != nil {
//			log.Fatal(err)
//		}
//		return cars
//	}
func GetCar(cl *mongo.Client, id primitive.ObjectID) interface{} {
	filter := bson.D{{"_id", id}}
	coll := cl.Database("data").Collection("cars")
	var car Car
	err := coll.FindOne(context.TODO(), filter).Decode(&car)
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}
func SetCar(cl *mongo.Client, doc Car) {
	coll := cl.Database("data").Collection("cars")
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateCar(cl *mongo.Client, id primitive.ObjectID, doc Car) {
	jsonDoc, _ := json.Marshal(doc)
	fmt.Println(string(jsonDoc))
	result := make(map[string]interface{})
	json.Unmarshal(jsonDoc, &result)
	coll := cl.Database("data").Collection("cars")
	update := bson.M{
		"$set": result,
	}
	_, err := coll.UpdateByID(context.TODO(), id, update)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteCar(cl *mongo.Client, id primitive.ObjectID) {
	coll := cl.Database("data").Collection("cars")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func GetProfile(cl *mongo.Client, id primitive.ObjectID) []byte {
	filter := bson.D{{"_id", id}}
	coll := cl.Database("data").Collection("profiles")
	var prof Profile
	err := coll.FindOne(context.TODO(), filter).Decode(&prof)
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(prof)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func SetProfile(cl *mongo.Client, doc Profile) {
	coll := cl.Database("data").Collection("profiles")
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateProfile(cl *mongo.Client, id primitive.ObjectID, doc Profile) {
	jsonDoc, _ := json.Marshal(doc)
	fmt.Println(string(jsonDoc))
	result := make(map[string]interface{})
	json.Unmarshal(jsonDoc, &result)
	coll := cl.Database("data").Collection("profiles")
	update := bson.M{
		"$set": result,
	}
	_, err := coll.UpdateByID(context.TODO(), id, update)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteProfile(cl *mongo.Client, id primitive.ObjectID) {
	coll := cl.Database("data").Collection("profiles")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func GetEvent(cl *mongo.Client, id primitive.ObjectID) interface{} {
	filter := bson.D{{"_id", id}}
	coll := cl.Database("data").Collection("events")
	var event Event
	err := coll.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}
func SetEvent(cl *mongo.Client, doc Event) {
	coll := cl.Database("data").Collection("events")
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateEvent(cl *mongo.Client, id primitive.ObjectID, doc Event) {
	jsonDoc, _ := json.Marshal(doc)
	fmt.Println(string(jsonDoc))
	result := make(map[string]interface{})
	json.Unmarshal(jsonDoc, &result)
	coll := cl.Database("data").Collection("events")
	update := bson.M{
		"$set": result,
	}
	_, err := coll.UpdateByID(context.TODO(), id, update)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteEvent(cl *mongo.Client, id primitive.ObjectID) {
	coll := cl.Database("data").Collection("events")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func InitiateMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func UploadFile(conn *mongo.Client, profileId, file string) {
	filename := profileId
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	//conn := InitiateMongoClient()
	bucket, err := gridfs.NewBucket(
		conn.Database("myfiles"),
	)
	if err != nil {
		log.Fatal(err)
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
}
func GetAvatar(conn *mongo.Client, file string) string {
	// For CRUD operations, here is an example
	db := conn.Database("myfiles")
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	// you can print out the results

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(file, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File size to download: %v\n", dStream)
	img := buf.Bytes()
	//data := make([]byte, base64.StdEncoding.EncodedLen(len(img)))
	return base64.StdEncoding.EncodeToString(img)
}

//func main() {
//	var client = InitiateMongoClient()
//	file := os.Args[1]
//	filename := path.Base(file)
//	UploadFile(client, file, filename)
//
//	//doc := Car{Make: "Toyota", Model: "Venza", Year: 2013,
//	//	Vin: "123qwe123qwe", PurchaseDate: "23-05-2014",
//	//	Transmission: "Auto", CurrentMileage: 123456,
//	//	BodyType: "Crossover"}
//	//update := bson.D{
//	//	{"$set", bson.D{{"model", "Venzaaaac"}}},
//	//}
//	//setCar(client, doc)
//	//objID, _ := primitive.ObjectIDFromHex("6405cd68794b31e58b02e89c")
//	//fmt.Println(objID)
//	//fmt.Println(getCars(client))
//	//updateCar(client, objID, update)
//	//deleteCar(client, objID)
//	//fmt.Println(getCars(client))
//
//	//fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
//
//	//fmt.Println(databases)
//}
