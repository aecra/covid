package main

type User struct {
	Name     string `db:"name"`
	Position string `db:"position"`
	State    string `db:"state"`
	Email    string `db:"email"`
	Eaisess  string `db:"eaisess"`
	Uukey    string `db:"uukey"`
}

type ClockRes struct {
	M string `json:"m"`
}
