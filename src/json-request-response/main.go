package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request defines the http request object
type Request struct {
	Name   string `json:"name"`
	Param1 int16  `json:"a"`
	Param2 int16  `json:"b"`
}

// Response defines the http response object
type Response struct {
	Status  int16  `json:"status"`
	Result  int16  `json:"result"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if r.Method == http.MethodPost {
			// Read json body
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error parsing request body", 400)
				return
			}

			// Prepare response
			var res Response
			res.Result = req.Param1 + req.Param2
			res.Message = "Addtion operation performed"
			res.Status = http.StatusOK

			// Send json response
			err = json.NewEncoder(w).Encode(res)
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
			}
		}
	})
	fmt.Println("Server Listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
