package utils

import "regexp"

// ValidatePhone is a func to Validate whether phone number is acceptable
func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^1(3|4|5|7|8)\d{9}$`)
	return re.MatchString(phone)
}
