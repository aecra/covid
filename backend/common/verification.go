package common

import "regexp"

func VerifyEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyHost(host string) bool {
	pattern := `^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(host)
}

func VerifyUsername(username string) bool {
	pattern := `^[a-zA-Z0-9_-]{5,16}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

func VerifyPassword(password string) bool {
	pattern := `^[a-zA-Z0-9!@#$%&*_-]{8,16}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}
