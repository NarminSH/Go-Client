package main

//we want to render our data as json
import (
	"encoding/json"
	"fmt"
	"strings"

	"clientapi/middleware"
	"clientapi/models"
	"log" //error handling
	"net/http"

	_ "clientapi/docs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux" // client/server functionality
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	httpSwagger "github.com/swaggo/http-swagger"
)





var db *gorm.DB
var err error

//slice named clients inherits from Client struct
var clients []models.Client




  type Succesful struct {
	// IsSuccess bool   `json:"isSucces"`
	Message string `json:"message"`
}


func SetSuccess(success Succesful, message string) Succesful {
	// success.IsSuccess = true
	success.Message = message
	return success
}


// GetClients godoc
// @Summary Get   all clients
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
// @Description get client by username
// @Tags Clients
// @Param username path string true "Client username"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{username} [get]
func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	var client models.Client

	user := db.Where("username = ?", username).First(&client)
	fmt.Println(user, "user over herereee", client.ID)
	json.NewEncoder(w).Encode(client)
}


// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client with the input paylod
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param newClient body models.Client true "Create client"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients [post]
// @Security ApiKeyAuth
func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("worked create")
	w.Header().Set("Content-Type", "application/json")
	var newClient models.Client
	if r.Header["Authorization"] == nil{
		var err middleware.Error
		err = middleware.SetError(err, "No Authorization credentials were provided")
		json.NewEncoder(w).Encode(err)
		return
	}

	stringToken := r.Header["Authorization"][0]
	split := strings.Split(stringToken, " ")
	responseToken := split[1]

	var mySigningKey = []byte(middleware.Secretkey)

	token, _ := jwt.Parse(responseToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing token.")
		}
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok{
			fmt.Println(claims, "claims are herererererer")
			username := claims["Username"] 
			StrUsername, _ := username.(string)
			json.NewDecoder(r.Body).Decode(&newClient)

			newClient.Username = StrUsername
			fmt.Println(StrUsername, "usernameee")
			
			// newClient.ID = strconv.Itoa(len(clients) + 1)
			// clients = append(clients, newClient)
			err := db.Where("username = ? ", StrUsername).First(&newClient).Error
			fmt.Println(err, "zeroooo")
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					fmt.Println(err, "firsttt")
					db.Create(&newClient)
				} 
			} else {
				fmt.Println(err, "seconddd")
				var err middleware.Error
				err = middleware.SetError(err, "User with this username exists")
				json.NewEncoder(w).Encode(err)
				return
			}

			json.NewEncoder(w).Encode((newClient))
}
}




// UpdateClient godoc
// @Summary Update particular client
// @Description Update particular client by username 
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param username path string true "Client username"
// @Param updatedclient body models.Client true "Update Client"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{username} [put]
// @Security ApiKeyAuth
func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestUser := params["username"]
	fmt.Println(requestUser, "requested user is over hereeeee")
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
// @Description Delete particular client by username
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param username path string true "Client username"
// @Success 200 {object} models.Client
// @Failure 400,404 {object} object
// @Router /clients/{username} [delete]
// @Security ApiKeyAuth
func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestUser := params["username"]
	fmt.Println(requestUser, "trying to delete user")
	// id64, _ := strconv.ParseUint(requestId, 10, 64)
	// idToDelete := uint(id64)
	db.Where("username = ?", requestUser).Delete(&models.Client{})
	
	w.WriteHeader(http.StatusOK)
	var success Succesful
	success = SetSuccess(success, "User is deleted!")
	json.NewEncoder(w).Encode(success)
	return
}




//get all orders of particular user 
func clientOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	var client models.Client
	user := db.Where("username = ?", username).First(&client)
	fmt.Println(user, "user over herereee")
	fmt.Println(client.ID, "client id is over here")
	client_id := client.ID
	var orders []models.Order
	db.Where("client_id = ? ", client_id).Find(&orders)
	json.NewEncoder(w).Encode(orders)
}



//get all ongoing orders of particular user 
func clientActiveOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	var client models.Client
	user := db.Where("username = ?", username).First(&client)
	fmt.Println(user, "user over herereee")
	fmt.Println(client.ID, "client id is over here")
	client_id := client.ID
	var orders []models.Order
	active := db.Where("client_id = ? AND complete = ? ", client_id, "False").Preload("Items").Find(&orders)
	fmt.Println(active, "active orderssss")
	json.NewEncoder(w).Encode(orders)
}



// Orders
// func createOrder(w http.ResponseWriter, r *http.Request) {
// 	var neworder models.Order
// 	json.NewDecoder(r.Body).Decode(&neworder)
// 	// Creates new order by inserting records in the `orders` and `items` table
// 	db.Create(&neworder)
// 	fmt.Println("crerated new order")
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




// host=client_postgres_1 - local - server
// host=localhost - local development
// host=192.168.31.74 - production

// @title Clients API
// @version 1.0
// @description This is a sample service for managing clients


//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization

// @host 192.168.31.74:8004
// @BasePath /api/v1.0
func main() {
	db, err = gorm.Open("postgres", "host=192.168.31.74  user=lezzetly password=lezzetly123 dbname=db_name port=5432 sslmode=disable Timezone=Asia/Baku")

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
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Item{})
	db.Model(&models.Order{}).AddForeignKey("client_id", "clients(id)", "NO ACTION", "NO ACTION")

	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/clients", getCLients).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{username}", getClient).Methods("GET")
	router.HandleFunc("/api/v1.0/clients", createClient).Methods("POST")
	router.HandleFunc("/api/v1.0/clients/{username}", middleware.IsAuthorized(updateClient)).Methods("PUT")
	router.HandleFunc("/api/v1.0/clients/{username}", middleware.IsAuthorized(deleteClient)).Methods("DELETE")
	router.HandleFunc("/api/v1.0/clients/{username}/orders", clientOrders).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{username}/active-orders", clientActiveOrders).Methods("GET")

	// router.HandleFunc("/api/v1.0/orders", getOrders).Methods("GET")
	// router.HandleFunc("/api/v1.0/orders", createOrder).Methods("POST")
	// router.HandleFunc("/api/v1.0/orders/{orderId}", getOrder).Methods("GET")
	// router.HandleFunc("/api/v1.0/orderitems", createItem).Methods("POST")


	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8000", router))


}
