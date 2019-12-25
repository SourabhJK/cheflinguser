package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	user "github.com/cheflinguser/chefling/usermanagement"
)

func Signup(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token")

	switch req.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	}
	var userInfo user.User

	signupBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("SignUp: Not Able to Read SignUp Request Body")
		return
	}

	err = json.Unmarshal(signupBody, &userInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("SignUp: Error in Unmarshalling SignUp Request Body")
		return
	}

	token, err := userInfo.SignUp()

	if err != nil {
		fmt.Println("Error in signup:", err)
		if err == user.UserExist {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("User already exist")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Please try again later")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func SignIn(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token")

	switch req.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	}
	var UserInfo user.User
	fmt.Println("Inside SignIn")
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR: Not Able to Read SignUp Request Body")
		return
	}

	err = json.Unmarshal(requestBody, &UserInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR: Error in Unmarshalling SignUp Request Body")
		return
	}

	token, err := UserInfo.SignIn()

	if err != nil {
		if err == user.UserNotFound {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Username or password mismatch")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Please try again later")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}

func Profile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token")

	switch req.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	}
	var (
		UserInfo user.User
		query    = req.URL.Query()
	)

	if query["username"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username is expected as query param")
	}

	UserInfo.Username = query["username"][0]
	UserInfo.Password = req.Header.Get("password")

	if UserInfo.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Password is expected in headers")
	}

	userDetails, err := UserInfo.Read()

	if err != nil {
		if err == user.UserNotFound {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Username or password mismatch")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Please try again later")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDetails)

}

func ProfileUpdate(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token")

	switch req.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	}
	var (
		userInfo    user.User
		query       = req.URL.Query()
		userDetails = make(map[string]interface{})
	)

	if query["user"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username is expected as query param")
	}

	userInfo.Username = query["user"][0]
	userInfo.Password = req.Header.Get("password")

	if userInfo.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Password is expected in headers")
	}

	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR: Not Able to Read SignUp Request Body")
		return
	}

	err = json.Unmarshal(requestBody, &userDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR: Error in Unmarshalling SignUp Request Body")
		return
	}

	err = userInfo.Update(userDetails)

	if err != nil {
		if err == user.UserNotFound {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Username or password mismatch")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Please try again later")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User updated successfully")

}
