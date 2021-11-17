package middleware

import (
	"clientapi/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var db *gorm.DB



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
	// IsError bool   `json:"isError"`
	Warning string `json:"warning"`
}


func SetError(err Error, message string) Error {
	// err.IsError = true
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

		token, err := jwt.Parse(responseToken, func(token *jwt.Token) (interface{}, error) {
			fmt.Println(token, "Parsed token is over here")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		// if err != nil {
		// 	var err Error
		// 	fmt.Println(err, "errrooorrrrr")
		// 	err = SetError(err, "Your Token has been expired.")
		// 	json.NewEncoder(w).Encode(err)
		// 	return
		// }

		if claims, ok := token.Claims.(jwt.MapClaims); ok{
			fmt.Println(r.Method, "methodddd")
			fmt.Println(claims, "claimsss are here")
			params := mux.Vars(r)
			requestUser := params["id"]
			fmt.Println(requestUser, "hhhhhhhhh")
			
			// if r.Method == "POST" {
			// 	username := claims["Username"] 
			// 	// Str, _ := username.(string)
			// 	fmt.Println("username type is", reflect.TypeOf(username))
			// 	handler.ServeHTTP(w, r)
			// 	return
			// }
			var client1 models.Client
			Claimed_user := claims["Username"]
			fmt.Println(Claimed_user, "qqqqqqqqq")
			StrUsername, _ := Claimed_user.(string)
			
			fmt.Println(StrUsername, "ssssssssss")
			fmt.Println(db, "dddddddd")
			// q := main.Where("username = ? ", StrUsername).First(&client1)
			// fmt.Println(q)
			// if err != nil {
			// 	if err == gorm.ErrRecordNotFound {
			// 		fmt.Println(err, "firsttt")
			// 	} 
			// } else{
			// 	fmt.Println("another error is here")
			// }
			// fmt.Println(s, "rrrrrrr")
			// fmt.Println(s, "query resultttttt")
			client_id := client1.ID
			var token_user string
			token_user = strconv.FormatUint(uint64(client_id), 10)
			fmt.Println(Claimed_user, "claimed user ")
			fmt.Println(token_user, "ccccccccc")


			if token_user == requestUser  {
				fmt.Println(token_user, "token user id iss ")
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



