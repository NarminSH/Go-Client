package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)


type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	UserType int `json:"user_usertype"`
	UserName string `json:"user_username"`
	IsFull bool `json:"user_isfull"`
}

var (
	router *mux.Router
	Secretkey string = "secretkeyjwt"
  )


type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}


func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
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
			requestUser := params["username"]

			if r.Method == "POST" {
				username := claims["iss"] 
				// Str, _ := username.(string)
				fmt.Println("username type is", reflect.TypeOf(username))
				handler.ServeHTTP(w, r)
				return
			}



			if claims["iss"] == requestUser  {
				fmt.Printf("User is %s ", requestUser)
				// r.Header.Set("iss", "nermin")
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
			}
		}

		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}
