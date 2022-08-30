package object

import (
	"errors"
	"math/rand"
	"time"

	"github.com/aecra/covid/db"
)

type User struct {
	Username string `gorm:"primaryKey" json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	State    bool   `json:"state"`
	Eaisess  string `json:"eaisess"`
	Uukey    string `json:"uukey"`
	Position string `json:"position"`
	Home     string `json:"home"`
}

func GetActiveClockUser() []User {
	var users []User
	db.GetConnection().Where(&User{State: true, Position: "school"}).Find(&users)
	// shuffle the users
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
	return users
}

func GetActiveHealthUser() []User {
	var users []User
	db.GetConnection().Where(&User{State: true, Position: "home"}).Find(&users)
	// shuffle the users
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
	return users
}

func GetUserByName(username string) (*User, error) {
	var user User
	db.GetConnection().Where("username = ?", username).Take(&user)
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func AddUser(user *User) {
	db.GetConnection().Create(user)
}

func UpdateUser(user *User) {
	db.GetConnection().Save(user)
}

func IsUniqueUser(username, email string) bool {
	var users []User
	db.GetConnection().Where("username = ? OR email = ?", username, email).Find(&users)
	return len(users) == 0
}

func Register(username, password, email string) error {
	if !IsUniqueUser(username, email) {
		return errors.New("username or email already exists")
	}
	user := User{Username: username, Password: password, Email: email, State: false}
	return db.GetConnection().Create(user).Error
}
