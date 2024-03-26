package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arvind910/momgo/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://arvi:arvi@cluster0.ezswte5.mongodb.net"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client Option
    clientOption := options.Client().ApplyURI(connectionstring)
   //connect to mongodb
    client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("MongoDB connection success")
	
	collection = client.Database(dbName).Collection(colName)

	//collection instance

	fmt.Println("Collection instance is ready")

}

// MongoDB helper -file
//insert 1 record

func insertOneMovie(movie model.Netflix) {
   inserted, err := collection.InsertOne(context.Background(), movie)
   if err != nil{
		    log.Fatal(err)
   }
   fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}


// update 1 record
func updateOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{ "$set": bson.M{"watched": true}}

result, err :=	collection.UpdateOne(context.Background(), filter, update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string)  {
    id, _ :=	primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
deleteCount, err :=	collection.DeleteOne(context.Background(), filter)
if err != nil{
	log.Fatal(err)
}
fmt.Println("delete count:",deleteCount )
}

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteCount, err:= collection.DeleteMany(context.Background(), filter)
	    if err != nil{
		log.Fatal(err)
		}
		fmt.Println("Number of movies deleted: ", deleteCount.DeletedCount)
		return deleteCount.DeletedCount
}
func getAllMovies() []primitive.M {
	
	cur, err := collection.Find(context.Background(), bson.D{{}})
	    if err != nil{
		log.Fatal(err)
	    }
		 defer cur.Close(context.Background())
	    var movies  []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)

	}
		movies = append(movies, movie)
	
}
     
    return movies

	
}

// Actual controller - file

func GetAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	
}
func DeleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	
}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
	
}