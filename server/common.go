package main

import (
	"github.com/sethvargo/go-password/password"
	"regexp"
)

const passwordRegex = `^(.{5,5})-(.{5,5})-(.{5,5})-(.{5,5})-(.{5,5})$`

func isValid(password string) bool {
	var re = regexp.MustCompile(passwordRegex)
	var matches = re.FindAllString(password, -1)

	if len(matches) == 0 {
		return false
	} else {
		return true
	}
}

func (p *Plugin) mask(password string) string {
	re := regexp.MustCompile(passwordRegex)
	string := re.ReplaceAllString(password, "XXXXX-XXXXX-XXXXX-XXXXX-$5")

	return string
}

func (p *Plugin) generatePassword() string {
	res, err := password.Generate(16, 2, 0, false, false)
	if err != nil {
		p.API.LogError("Could not generate password")
	}

	return res
}
