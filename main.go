package main

import (
	"encoding/json"
	"fmt"
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", registerHandler)

	println("localhost:8080 is running")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"message":"every thing is ok "}`)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, `{"error":"invalid method "}`)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	var req userservice.RegisterRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Register(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}
	w.Write([]byte(`{"message":"user created"}`))
}
