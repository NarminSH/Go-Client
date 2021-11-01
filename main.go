package main

//we want to render our data as json
import (
	"encoding/json"
	"fmt"

	"log" //error handling
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux" // client/server functionality
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Client struct {
	ID          uint   `json:"id" gorm:"primary_key"` //properties
	FirstName   string `json:"firstname" gorm:"type:varchar(50)"`
	LastName    string `json:"lastname" gorm:"type:varchar(50)"`
	Patronymic  string `json:"patronymic" gorm:"type:varchar(70)"`
	Username    string `json:"username" gorm:"type:varchar(70)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(70)"`
	Email       string `json:"email" gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//slice named clients inherits from Client struct
var clients []Client

func getCLients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Find(&clients)
	json.NewEncoder(w).Encode(clients)

}

func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var client Client
	db.First(&client, id)
	json.NewEncoder(w).Encode(client)
}

// http.ResponseWriter -  contains the response details(headers, payload)
// http.Request -  contains the incoming request details

func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("worked create")
	w.Header().Set("Content-Type", "application/json")
	var newClient Client
	json.NewDecoder(r.Body).Decode(&newClient)
	// newClient.ID = strconv.Itoa(len(clients) + 1)
	// clients = append(clients, newClient)
	db.Create(&newClient)
	json.NewEncoder(w).Encode((newClient))
}


func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestId := params["id"]
	id64, _ := strconv.ParseUint(requestId, 10, 64)
	idToUpdate := uint(id64)
	var updatedclient Client
	json.NewDecoder(r.Body).Decode(&updatedclient)
	updatedclient.ID = idToUpdate
	db.Where("id = ?", idToUpdate).Save(&updatedclient)
	json.NewEncoder(w).Encode(updatedclient)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestId := params["id"]
	id64, _ := strconv.ParseUint(requestId, 10, 64)
	idToDelete := uint(id64)
	db.Where("id = ?", idToDelete).Delete(&Client{})
	w.WriteHeader(http.StatusOK)
}


// host=client_postgres_1 - local - server
// host=localhost - local development
// host=192.168.31.74 - production
func main() {
	db, err = gorm.Open("postgres", "host=192.168.31.74 user=lezzetly password=lezzetly123 dbname=db_name port=5432 sslmode=disable Timezone=Asia/Baku")

	if err != nil {
		fmt.Println(err, "error is  here")
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	// db.Exec("CREATE DATABASE client_db")
	// db.Exec("USE client_db")
	db.AutoMigrate(&Client{})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/clients", getCLients).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{id}", getClient).Methods("GET")
	router.HandleFunc("/api/v1.0/clients", createClient).Methods("POST")
	router.HandleFunc("/api/v1.0/clients/{id}", updateClient).Methods("PUT")
	router.HandleFunc("/api/v1.0/clients/{id}", deleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
