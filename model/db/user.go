package db

type User struct {
	Id           string `bson:"_id,omitempty"`
	UserName     string `bson:"username"`
	Email        string `bson:"email"`
	HashPassword []byte `bson:"hashpassword"`
}
