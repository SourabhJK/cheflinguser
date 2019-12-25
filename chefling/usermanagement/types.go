package usermanagement

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	EmailId string  `json:"emailid"`
}