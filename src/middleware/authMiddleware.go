package middleware

import (
  "fmt"
  "net/http"
  "github.com/dgrijalva/jwt-go"
  "github.com/dgrijalva/jwt-go/request"
)

var VerifyKey = []byte("arg98634hkjmngerh45eh847e36h514aher5hh3erfhs3e1fhwehw4eh1d3654ju84kgh1ndfg3n1st84hw9te87rtyd6h1sd6fg1ne5rt46uet4yui")

func Authorization(next http.HandlerFunc) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
            func(token *jwt.Token) (interface{}, error) {
    		       return VerifyKey, nil
  	   })

    if err == nil {
    	if token.Valid {
    		next(w, r)
    	} else {
    		w.WriteHeader(http.StatusUnauthorized)
    		fmt.Fprint(w, "Token is not valid")
    	}
    } else {
    	w.WriteHeader(http.StatusUnauthorized)
    	fmt.Fprint(w, "Unauthorised access to this resource")
    }
  })
}
