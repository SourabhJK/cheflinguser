package usermanagement

type User struct {
	Id       uint64
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	EmailId  string `json:"emailid" db:"emailid"`
}
