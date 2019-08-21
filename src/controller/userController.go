package controller

import(
    "net/http"
    "../model"
    //"strconv"
    "encoding/json"
  )

func GetUserInformation(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var data model.User
    if err := decoder.Decode(&data); err != nil{
        panic(err)
    }

    model.Insert(data)
    users := model.FetchAll()

    w.Header().Set("Content-Type" , "application/json; charset=UTF-8")
    w.WriteHeader(200)
    if err := json.NewEncoder(w).Encode(users); err != nil{
        panic(err)
    }
}
