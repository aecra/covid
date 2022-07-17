package object

type User struct {
	Name     string `gorm:"primaryKey" json:"name" binding:"required"`
	Position string `json:"position"`
	State    bool   `json:"state"`
	Email    string `json:"email"`
	Eaisess  string `json:"eaisess"`
	Uukey    string `json:"uukey"`
	Home     string `json:"home"`
}

func GetActiveUser() []User {
	var users []User
	database.Where(&User{State: true}).Find(&users)
	return users
}

func GetUserByName(name string) User {
	var user User
	database.Where("name = ?", name).Take(&user)
	return user
}

func AddUser(user *User) {
	database.Create(user)
}

func UpdateUser(user *User) {
	database.Save(user)
}
