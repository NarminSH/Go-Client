package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	secretkey string = "secretkeyjwt"
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

		var mySigningKey = []byte(secretkey)

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


		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
			fmt.Println(claims, "claimsss are here")
			params := mux.Vars(r)
			requestUser := params["username"]
			if claims["iss"] == requestUser  {
				fmt.Printf("User is %s ", requestUser)
				// r.Header.Set("iss", "nermin")
				handler.ServeHTTP(w, r)
				return
			}else{
				fmt.Println("User is not same with updated user")
				var err Error
				err = SetError(err, "You can not update the client")
				json.NewEncoder(w).Encode(err)
				return
			}
		}

		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}
