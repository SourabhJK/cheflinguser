package usermanagement

import(
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	log "github.com/Sirupsen/logrus"
	"errors"
	"time"
	"github.com/cheflinguser/db"
	"encoding/json"
	"fmt"
)


const (
	JWT_KEY = "secret"
	UserCollection = "users"
	
)


var (
	UserExist = errors.New("User Already exist")
	UserNotFound = errors.New("User Not found")
)



func (user User)Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user User) GenerateUserToken() (string, error) {

	ExpiresAt := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      float64(ExpiresAt),
	})
	return token.SignedString([]byte(JWT_KEY))

}


func (user User) SignUp() (string, error){

	err := user.checkuser()

	if err != nil {
		log.Error("Error in SignUp --> checkuser: "+ err.Error())
		return "",err
	}

	err = user.encryptPassword()

	if err != nil {
		log.Error("Error in SignUp --> encryptPassword: "+ err.Error())
		return "", err
	}


	err = user.Create()
	if err != nil {
		log.Error("Error in SignUp --> Create: "+ err.Error())
		return "", err
	}

	token, err := user.GenerateUserToken()

	if err != nil{
		log.Error("Error in SignUp --> GenerateUserToken: "+ err.Error())
	}
	
	return token, err

}

func (user User) SignIn() (string, error){
	var (
		token string
		err error
	)
	
	user, err = user.ReadOne()
	if err != nil{
		log.Error("Error in SignIn --> Read: "+ err.Error())
		return token, err
	}
	
	token, err = user.GenerateUserToken()

	if err != nil{
		log.Error("Error in SignIn --> GenerateUserToken: "+ err.Error())
	}
	
	return token, err

}

func (user User) Create() error{
	fmt.Println("Inside create()")
	return db.DBOpt.DB.Create(UserCollection, user)

}

func (user User) Read() ([]User, error){
	var output []User
	findQuery, err := user.convert()
	if err != nil{
		log.Error("Error in Read --> convert:", err.Error())
	}
	result, err := db.DBOpt.DB.Read(UserCollection, findQuery)
	if err != nil{
		log.Error("Error in Read --> Read:"+ err.Error())
	}

	if result == nil{
		return output, UserNotFound
	}

	output, _ = result.([]User)
	
	return output,nil
}

func (user User) ReadOne() (User, error){
	var output User
	findQuery, err := user.convert()
	if err != nil{
		log.Error("Error in Read --> convert:", err.Error())
	}
	result, err := db.DBOpt.DB.Read(UserCollection, findQuery)
	if err != nil{
		log.Error("Error in Read --> Read:"+ err.Error())
	}

	if result == nil{
		return output, UserNotFound
	}
	output, _ = result.(User)
	return output, err
}

func (user User) Update(updateMap map[string]interface{}) (error){
	findQuery, err := user.convert()
	if err != nil{
		log.Error("Error in Read --> convert:", err.Error())
	}
	err = db.DBOpt.DB.Update(UserCollection, findQuery, updateMap)
	if err != nil{
		log.Error("Error in Update --> Update:", err.Error())
	}
	return err
}


func (user User) checkuser() error{
	return nil
}



func (user *User) encryptPassword() error{
	var err error
	user.Password, err = HashPassword(user.Password)
	return err
}


func (user User) convert() (map[string]interface{}, error){
	var output = make(map[string]interface{})
	data, err := json.Marshal(user)

	if err != nil{
		log.Error("Error in convert --> Marshal: "+ err.Error())
		return output, err
	}

	err = json.Unmarshal(data, &output)
	if err != nil{
		log.Error("Error in convert --> Marshal: "+ err.Error())
	}

	return output, err
}



func HashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}


