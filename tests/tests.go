package tests




// func TestGetRequest(t *testing.T) {
//     t.Parallel()

//     r, _ := http.NewRequest("GET", "/test/abcd", nil)
//     w := httptest.NewRecorder()

//     //Hack to try to fake gorilla/mux vars
//     vars := map[string]string{
//         "mystring": "abcd",
//     }

//     // CHANGE THIS LINE!!!
//     r = mux.SetURLVars(r, vars)

//     GetRequest(w, r)

//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.Equal(t, []byte("abcd"), w.Body.Bytes())
// }