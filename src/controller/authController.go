package controller

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "github.com/dgrijalva/jwt-go"
    "../model"
    "time"
)

var SignKey = []byte("arg98634hkjmngerh45eh847e36h514aher5hh3erfhs3e1fhwehw4eh1d3654ju84kgh1ndfg3n1st84hw9te87rtyd6h1sd6fg1ne5rt46uet4yui")

type Token struct {
    Access_Token 	string    `json:"access_token"`
}

func Login(w http.ResponseWriter, r *http.Request){
    decoder := json.NewDecoder(r.Body)
    var loginInfo model.User

    if err := decoder.Decode(&loginInfo); err != nil{
        panic(err)
    }

    if loginInfo.Username == "admin" && loginInfo.Password == "admin"{
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "iss": "admin",
            "exp": time.Now().Add(time.Minute * 10).Unix(),
            "CustomUserInfo": struct {
                 Name string
                 Role string
            }{loginInfo.Username, "Member"},})

        tokenString, err := token.SignedString(SignKey)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintln(w, "Error while signing the token")
            log.Printf("Error signing token: %v\n", err)
        }

        //create a token instance using the token string
        response := Token{tokenString}

        w.Header().Set("Content-Type" , "application/json; charset=UTF-8")
        w.WriteHeader(200)
        if err := json.NewEncoder(w).Encode(response); err != nil{
            panic(err)
        }
    }else{
        w.WriteHeader(http.StatusUnauthorized)
    }
}
