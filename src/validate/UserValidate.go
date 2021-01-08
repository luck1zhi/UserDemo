package validate

import (
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func CheckUserId(fields validator2.FieldLevel) bool {
	pattern := "\\d+"
	if userId, ok := fields.Field().Interface().(string); ok {
		if bool, _ := regexp.MatchString(pattern, userId); bool {
			return false
		}
	}
	return true
}

func CheckUserName(fields validator2.FieldLevel) bool {
	fmt.Println("fields validator2.FieldLevel")
	if userName, ok := fields.Field().Interface().(string); ok {
		if strings.TrimSpace(userName) != "" && len(userName) >= 6 && len(userName) <=12 {
			return false
		}
	}
	return true
}