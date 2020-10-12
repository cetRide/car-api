package validation

import (
	"regexp"
)

func ValidatePassword(password string) bool {
	// password contains a capital letter
	r, _ := regexp.Compile(`[A-Z]`)
	if !r.MatchString(password) {
		return false
	}
	//  password contain lowercase letter
	r, _ = regexp.Compile(`[a-z]`)
	if !r.MatchString(password) {
		return false
	}
	// password contain number
	r, _ = regexp.Compile(`[0-9]`)
	if !r.MatchString(password) {
		return false
	}
	if len(password) < 8 {
		return false
	}

	return true
}

//email validation
func ValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return false
	}
	return true

}
