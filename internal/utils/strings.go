package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail checks if a string is a valid email address
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	emailRegex := regexp.MustCompile(`^[\w+\-.]+@[a-z\d\-.]+\.[a-z]+$`)
	return emailRegex.MatchString(email)
}

// IsValidPhoneNumber checks if a string is a valid phone number
func IsValidPhoneNumber(phoneNumber string) bool {
	if phoneNumber == "" {
		return false
	}
	phoneRegex := regexp.MustCompile(`^[\d\+\-]+$`) // Regex to allow only numbers, plus, and dash
	return phoneRegex.MatchString(phoneNumber)
}

// Convert name into snake_case
func ToSnakeCase(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")
	return name
}
