package golash

import "regexp"

func IsEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email) == true
}

func IsName(name string) bool {
	re := regexp.MustCompile("(^[a-zA-Z\\s]+$)")
	return re.MatchString(name) == true
}
