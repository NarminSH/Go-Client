package main

//we want to render our data as json
import (
	"context"
	"encoding/json"
	"fmt"

	httpCheck "github.com/hellofresh/health-go/v4/checks/http"

	"strconv"
	"strings"
	"time"

	"github.com/alexliesenfeld/health"

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


type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	UserType int `json:"user_usertype"`
	UserName string `json:"user_username"`
	IsFull bool `json:"user_isfull"`
}

var (
	router *mux.Router
	Secretkey string = "supersecretkey"
)

var Claimed_user string


type Error struct {
	Warning string `json:"warning"`
}


func SetError(err Error, message string) Error {
	err.Warning = message
	return err
}
	


func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil{
			var err Error
			err = SetError(err, "No Authorization credentials were provided")
			json.NewEncoder(w).Encode(err)
			return
		}

		stringToken := r.Header["Authorization"][0]
		split :=strings.Split(stringToken, " ")
		responseToken := split[1]

		var mySigningKey = []byte(Secretkey)

		token, invalid := jwt.Parse(responseToken, func(token *jwt.Token) (interface{}, error) {
			fmt.Println(token, "Parsed token is over here")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})
		if invalid != nil {
			var err middleware.Error
				err = middleware.SetError(err, "Token is invalid")
				fmt.Println(invalid, "invalid token")
				json.NewEncoder(w).Encode(err)
				return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok{
			params := mux.Vars(r)
			requestUser := params["id"]
		
		
			var client models.Client
			Claimed_user := claims["Username"]
			fmt.Println(Claimed_user, "claimed user")
			
			err := db.Where("username = ? ", Claimed_user).First(&client).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					var err middleware.Error
					err = middleware.SetError(err, "Record was not found")
					json.NewEncoder(w).Encode(err)
					return
				} 
			}
			fmt.Println(client.ID, "found client id in db")
			// if err != nil {
			// 	if err == gorm.ErrRecordNotFound {
			// 		fmt.Println(err, "firsttt")
			// 	} 
			// } else{
			// 	fmt.Println("another error is here")
			// }
			// fmt.Println(s, "rrrrrrr")
			// fmt.Println(s, "query resultttttt")
			client_id := client.ID
			// fmt.Println(client_id, "fffffffffffffff")
			var tokenUser string
			tokenUser = strconv.FormatUint(uint64(client_id), 10)
			fmt.Println(tokenUser, "ccccccccc")


			if tokenUser == requestUser  {
					// if r.Method == "PUT" {
					// 	username := claims["Username"] 
					// 	StringUsername, _ := username.(string)
					// 	var updatedclient models.Client
					// 	updatedclient.Username = StringUsername
					// 	fmt.Println("username type is", reflect.TypeOf(username))
					// 	handler.ServeHTTP(w, r)
					// 	return
					// }
				fmt.Printf("User is %s ", requestUser)
				// r.Header.Set("Username", "nermin")
				handler.ServeHTTP(w, r)
				return
			}else {
				if r.Method == "PUT" {
				fmt.Println("User is not same with updated user")
				fmt.Printf("User is %s ", requestUser)
				var err Error
				err = SetError(err, "Only Client himself can change his profile")
				json.NewEncoder(w).Encode(err)
				return}
				if r.Method == "DELETE" {
					var err Error
					err = SetError(err, "Only Client himself can delete his profile")
					json.NewEncoder(w).Encode(err)
					return
				}
				if r.Method == "GET" {
					var err Error
					err = SetError(err, "You do not own permissions to look at others' profile")
					json.NewEncoder(w).Encode(err)
					return
				}
			}
		}

		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}




//slice named clients inherits from Client struct
var clients []models.Client


  type Succesful struct {
	Message string `json:"message"`
}


func SetSuccess(success Succesful, message string) Succesful {
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
	user_id := params["id"]
	fmt.Println(user_id, "get client user id")
	var client models.Client
	// fmt.Println(db, "dbdbdbdbdbdbdbbdbdddddddddd")
	user := db.Where("id = ?", user_id).First(&client)
	fmt.Println(user, "user id over herereee", client.ID)
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

	token, invalid := jwt.Parse(responseToken, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("entered token section")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("entered token parsing section")
			return nil, fmt.Errorf("There was an error in parsing token.")
		}
		return mySigningKey, nil
	})
	if invalid != nil {
		var err middleware.Error
			err = middleware.SetError(err, "Token is invalid")
			json.NewEncoder(w).Encode(err)
			return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok{
			fmt.Println(claims, "claims are hereeeeeeeee")
			username := claims["Username"] 
			usertype := claims["Usertype"]
			StrUsername, _ := username.(string)
			StrUsertype, _ := usertype.(string)
			if StrUsertype == "3" {
				json.NewDecoder(r.Body).Decode(&newClient)

				newClient.Username = StrUsername
				newClient.UserType = StrUsertype
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
			} else {
					fmt.Println(err, "another type user is trying to create a client")
					var err middleware.Error
					err = middleware.SetError(err, "You can not create client with this token")
					json.NewEncoder(w).Encode(err)
					return
			} 
			
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
		username := claims["Username"] 
		StringUsername, _ := username.(string)
		usertype := claims["Usertype"]
		StrUsertype, _ := usertype.(string)
		params := mux.Vars(r)
		requestUser := params["id"]
		id64, err := strconv.ParseUint(requestUser, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		Userid := uint(id64)

		var updatedclient models.Client
		json.NewDecoder(r.Body).Decode(&updatedclient)
		updatedclient.ID = Userid
		updatedclient.Username = StringUsername
		updatedclient.UserType = StrUsertype
		db.Where("id = ?", Userid).Save(&updatedclient)
		json.NewEncoder(w).Encode(updatedclient)
}
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
	requestUser := params["id"]
	fmt.Println(requestUser, "trying to delete user")
	id64, _ := strconv.ParseUint(requestUser, 10, 64)
	idToDelete := uint(id64)
	db.Where("id = ?", idToDelete).Delete(&models.Client{})
	
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
	user_id := params["id"]
	var client models.Client
	user := db.Where("id = ?", user_id).First(&client)
		fmt.Println(user, "user over herereee")
		fmt.Println(client.ID, "client id is over here")
		client_id := client.ID
		var orders []models.Order
		db.Where("client_id = ? ", client_id).Preload("Items").Find(&orders)
		json.NewEncoder(w).Encode(orders)
}
	




//get all ongoing orders of particular user 
func clientActiveOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestUser := params["id"]
	var client models.Client
	user := db.Where("id = ?", requestUser).First(&client)
	fmt.Println(user, "user over herereee")
	fmt.Println(client.ID, "client id is over here")
	client_id := client.ID
	var orders models.Order
	db.Where("client_id = ? AND complete = ? ", client_id, "False").Preload("Items").Find(&orders)
	fmt.Println(len(orders.Items), "Items are")
	json.NewEncoder(w).Encode(orders)
}



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
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	sqlDB := db.DB()
	fmt.Printf("database type: %T\n", sqlDB)

	err = sqlDB.Ping()
	if err != nil {
	panic(err)
	} 
	
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	// db.Exec("CREATE DATABASE client_db")
	// db.Exec("USE client_db")
	db.AutoMigrate(&models.Client{})
	// db.Set("gorm:auto_preload", true)
	// db.AutoMigrate(&models.Order{})
	// db.AutoMigrate(&models.Cook{})
	// db.AutoMigrate(&models.Courier{})
	// db.Model(&models.Order{}).AddForeignKey("client_id", "clients(id)", "NO ACTION", "NO ACTION")

	router := mux.NewRouter()


	checker := health.NewChecker(
		health.WithCacheDuration(10*time.Second),

		health.WithCheck(health.Check{
			Name:    "getClient",
			Check:   httpCheck.New(httpCheck.Config{
			   URL: "http://192.168.31.74:8004/api/v1.0/clients",
			}),
		 }),

		 health.WithCheck(health.Check{
			Name:    "getCook",
			Check:   httpCheck.New(httpCheck.Config{
			   URL: "http://192.168.31.74/api/v1.0/cooks",
			}),
		 }),


		health.WithCheck(health.Check{
			Name:    "database",      // A unique check name.
			Timeout: 5 * time.Second, // A check specific timeout.
			Check:   sqlDB.PingContext,
		}),
		// health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
		// 	Name:               "get",
		// 	Check:              sqlDB.PingContext,
		// 	StatusListener: func(ctx context.Context, name string, state health.CheckState) {
		// 	},
		// 	Interceptors: []health.Interceptor{},
		// }),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Println(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)

	router.HandleFunc("/health", health.NewHandler(checker))
	router.HandleFunc("/api/v1.0/clients", getCLients).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{id}", getClient ).Methods("GET")
	router.HandleFunc("/api/v1.0/clients", createClient).Methods("POST")
	router.HandleFunc("/api/v1.0/clients/{id}", IsAuthorized(updateClient)).Methods("PUT")
	router.HandleFunc("/api/v1.0/clients/{id}", IsAuthorized(deleteClient)).Methods("DELETE")
	router.HandleFunc("/api/v1.0/clients/{id}/orders", IsAuthorized(clientOrders)).Methods("GET")
	router.HandleFunc("/api/v1.0/clients/{id}/active-orders", IsAuthorized(clientActiveOrders)).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}


