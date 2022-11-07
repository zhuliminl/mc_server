package helper

import "regexp"

func IsPhoneValid(phone string) bool {
	return true
}
func IsPasswordValid(phone string) bool {
	return true
}
func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
