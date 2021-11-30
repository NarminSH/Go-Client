package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGetRequest(t *testing.T) {
//     router := mux.NewRouter() //initialise the router
//     testServer := httptest.NewServer(router) //setup the testing server
//     fmt.Println(testServer)
//      request,_ := http.NewRequest("GET","/api/v1.0/clients", nil)
//         resp := httptest.NewRecorder()
//         handler := http.HandlerFunc(getCLients)
// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// // directly and pass in our Request and ResponseRecorder.
//         handler.ServeHTTP(resp, request)
// }
func TestGetCLients(t *testing.T) {
	r := InitRouter()

	req, _ := http.NewRequest("GET", "/api/v1.0/clients", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
    db.Find(&clients)
    fmt.Println("client", clients)
    
	want, _ := json.Marshal(clients)
    fmt.Printf("type of clients is %T \n", want)
    string_want := string(want)
	got := res.Body.String()
    
    fmt.Printf("bgwbahxniwbhscxniuwhcnwciuhjwcb %T \n", string_want) 
    fmt.Printf("gotttt %T \n", got)

	// if got != string_want {
	// 	t.Errorf("want %v; got %v", string_want, got)
	// }
    assert.Equal(t, string_want, got, "The two answers should be the same.")
}