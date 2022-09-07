package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	paper "github.com/ersa97/paper-test/controllers"
	"github.com/ersa97/paper-test/middleware"
	"github.com/ersa97/paper-test/models"
	"github.com/gorilla/mux"
)

func Mux(paper paper.PaperService) {
	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})

	r.HandleFunc("/", paper.TestAPI).Methods("GET")
	//register
	r.HandleFunc("/register", paper.Register).Methods("POST")
	//login
	r.HandleFunc("/login", paper.Login).Methods("POST")
	//logout
	r.HandleFunc("/logout", middleware.Auth(paper.Logout)).Methods("GET")

	//financeaccount
	r.HandleFunc("/account/new", middleware.Auth(paper.AddAccount)).Methods("POST")
	r.HandleFunc("/account/{code}", middleware.Auth(paper.GetDetailAccount)).Methods("GET")
	r.HandleFunc("/account", middleware.Auth(paper.GetAccountList)).Methods("GET")
	r.HandleFunc("/account/{code}", middleware.Auth(paper.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/account/{code}", middleware.Auth(paper.DeleteAccount)).Methods("DELETE")

	//transaction
	r.HandleFunc("/transaction/new", middleware.Auth(paper.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{trxid}", middleware.Auth(paper.GetDetailTransaction)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(paper.GetListTransaction)).Methods("GET")
	r.HandleFunc("/transaction/{trxid}", middleware.Auth(paper.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/transaction/{trxid}", middleware.Auth(paper.DeleteTransaction)).Methods("DELETE")

	r.Use(mux.CORSMethodMiddleware(r))

	r.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "route not found",
			Data:    nil,
		})
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(models.Response{
			Message: "method not allowed",
			Data:    nil,
		})

	})

	appPort := os.Getenv("APPLICATION_PORT")

	log.Println("Running at " + os.Getenv("APP_URL") + ":" + appPort + "/")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", appPort), r)

}
