package router

import(
	"net/http"
	"github.com/cheflinguser/handler"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"github.com/cheflinguser/config"
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"fmt"
)

const JWT_KEY="secret"


func Route(){
	fmt.Println("Route....")
	r := mux.NewRouter()

	router := r.NewRoute().Subrouter()
	router.Use(ValidateMiddleware)

	//this route will be used for instancegroup healthcheck
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This is a catch-all route")
	})

	
	router.HandleFunc("/user/signup", handler.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/signin", handler.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/profile", handler.Profile).Methods("GET", "OPTIONS") 
	//Can be kept as /user/profile with the method as PUT
	router.HandleFunc("/user/profile/update", handler.ProfileUpdate).Methods("PUT", "OPTIONS") 

	http.ListenAndServe(":8090", handlers.LoggingHandler(os.Stdout, r))
	fmt.Println("Router started and listening at port 8090")

}

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token,test")
		switch req.Method {
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("Inside ValidateMiddleware.....")
		testToken := config.GetConfig().TestToken

		authorizationHeader := req.Header.Get("Authorization")
		

		if authorizationHeader == ""{
			json.NewEncoder(w).Encode("An authorization header is required")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if authorizationHeader == testToken{
			next.ServeHTTP(w, req)
			return
		}

		

			token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(JWT_KEY), nil
			})
			if error != nil {

				if error.Error() == "Token is expired" {
					w.WriteHeader(498)
				}

				json.NewEncoder(w).Encode("Token is expired. Login one more time")
				return
			}
			if token.Valid {
				context.Set(req, "decoded", token.Claims)
				next.ServeHTTP(w, req)
				return
			} else {
				json.NewEncoder(w).Encode("Invalid authorization token")
				w.WriteHeader(498)
				return
			}

		
	})
}