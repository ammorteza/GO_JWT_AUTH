package router

import "fmt"
import "log"
import "net/http"
import "../controller"
import "../middleware"
import "github.com/gorilla/mux"

func Route(){
    router := mux.NewRouter().StrictSlash(true)
    router.Handle("/user/insert", middleware.Authorization(http.HandlerFunc(controller.GetUserInformation))).Methods("POST")
    router.HandleFunc("/login", controller.Login).Methods("POST")

    fmt.Println("router is Listening to 8080 port number")
    log.Fatal(http.ListenAndServe(":8080", router))
}
