package main

import (
	"regexp"
)

type Registry struct {
	FullName string `json:"FullName"`
	Email    string `json:"Email"`
	Pass     string `json:"Pass"`
	PassConf string `json:"PassConf"`
	Errors   map[string]string
}

func ValidateEmail(registry *Registry) bool {
	registry.Errors = make(map[string]string)
	re := regexp.MustCompile(".+@-+\\..+")
	matched := re.Match([]byte(registry.Email))
	if matched == false {
		registry.Errors["Email"] = "Please enter a valid email address"
	}
	return len(registry.Errors) == 0

}

func main() {

}
