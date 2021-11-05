package main

//we want to render our data as json
import (
	"encoding/json"
	"fmt"

	"clientapi/middleware"
	"clientapi/models"
	"log" //error handling
	"net/http"

	_ "clientapi/docs"

	"github.com/gorilla/mux" // client/server functionality
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	httpSwagger "github.com/swaggo/http-swagger"
)





var db *gorm.DB
var err error

//slice named clients inherits from Client struct
var clients []models.Client




// GetClients godoc
// @Summary Get  all clients
// @Description Get  all clients
// @Tags Clients
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Client
// @Router /clients [get]
func getCLients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Find(&clients)
	json.NewEncoder(w).Encode(clients)

}

// @Summary Get one client
// @Description get client by ID
// @Tags Clients
// @Param id path string false "Client ID"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{id} [get]
func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var client models.Client
	db.First(&client, id)
	json.NewEncoder(w).Encode(client)
}


// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client with the input paylod
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param newClient body models.Client false "Create client"
// @Success 200 {object} models.Client
// @Router /clients [post]
func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("worked create")
	w.Header().Set("Content-Type", "application/json")
	var newClient models.Client
	json.NewDecoder(r.Body).Decode(&newClient)
	// newClient.ID = strconv.Itoa(len(clients) + 1)
	// clients = append(clients, newClient)
	db.Create(&newClient)
	json.NewEncoder(w).Encode((newClient))
}




// UpdateClient godoc
// @Summary Update particular client
// @Description Update particular client by id
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param updatedclient body models.Client false "Update Client"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{id} [put]
func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestUser := params["username"]
	// id64, _ := strconv.ParseUint(requestId, 10, 64)
	// idToUpdate := uint(id64)
	var updatedclient models.Client
	json.NewDecoder(r.Body).Decode(&updatedclient)
	updatedclient.Username = requestUser
	db.Where("username = ?", requestUser).Save(&updatedclient)
	json.NewEncoder(w).Encode(updatedclient)
}




// DeleteClient godoc
// @Summary Delete particular client
// @Description Delete particular client by id
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param idToDelete body models.Client false "Delete Client"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{id} [delete]
func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestUser := params["username"]
	// id64, _ := strconv.ParseUint(requestId, 10, 64)
	// idToDelete := uint(id64)
	db.Where("username = ?", requestUser).Delete(&models.Client{})
	w.WriteHeader(http.StatusOK)
}



// // Orders
// func createOrder(w http.ResponseWriter, r *http.Request) {
// 	var neworder models.Order
// 	json.NewDecoder(r.Body).Decode(&neworder)
// 	// Creates new order by inserting records in the `orders` and `items` table
// 	db.Create(&neworder)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(neworder)
// }


// func getOrders(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var orders []models.Order
// 	db.Preload("Items").Find(&orders)
// 	json.NewEncoder(w).Encode(orders)
// }



// func getOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	inputOrderID := params["orderId"]

// 	var order models.Order
// 	db.Preload("Items").First(&order, inputOrderID)
// 	json.NewEncoder(w).Encode(order)
// }

// // func updateOrder(w http.ResponseWriter, r *http.Request) {
// // 	var updatedOrder models.Order
// // 	json.NewDecoder(r.Body).Decode(&updatedOrder)
// // 	db.Save(&updatedOrder)

// // 	w.Header().Set("Content-Type", "application/json")
// // 	json.NewEncoder(w).Encode(updatedOrder)
// // }

// // Orders
// func createItem(w http.ResponseWriter, r *http.Request) {
// 	var newItem models.Item
// 	json.NewDecoder(r.Body).Decode(&newItem)
// 	// Creates new order by inserting records in the `orders` and `items` table
// 	db.Create(&newItem)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(newItem)
// }











// host=client_postgres_1 - local - server
// host=localhost - local development
// host=192.168.31.74 - production

// @title Clients API
// @version 1.0
// @description This is a sample serice for managing clients
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /
func main() {
	db, err = gorm.Open("postgres", "host=localhost user=lezzetly password=lezzetly123 dbname=db_name port=5432 sslmode=disable Timezone=Asia/Baku")

	if err != nil {
		fmt.Println(err, "Error is  here")
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	// db.Exec("CREATE DATABASE client_db")
	// db.Exec("USE client_db")
	db.AutoMigrate(&models.Client{})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/clients", getCLients).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{id}", getClient).Methods("GET")
	router.HandleFunc("/api/v1.0/clients", createClient).Methods("POST")
	router.HandleFunc("/api/v1.0/clients/{username}", middleware.IsAuthorized(updateClient)).Methods("PUT")
	router.HandleFunc("/api/v1.0/clients/{id}", middleware.IsAuthorized(deleteClient)).Methods("DELETE")
	// router.HandleFunc("/api/v1.0/orders", getOrders).Methods("GET")
	// router.HandleFunc("/api/v1.0/orders", createOrder).Methods("POST")
	// router.HandleFunc("/api/v1.0/orders/{orderId}", getOrder).Methods("GET")
	// router.HandleFunc("/api/v1.0/orderitems", createItem).Methods("POST")


	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8000", router))


}
