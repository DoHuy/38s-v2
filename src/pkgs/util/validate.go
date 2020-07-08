package util

import (
	"regexp"
	"strings"
)

var (
	phonePattern = regexp.MustCompile(`(^0\d{9}$)|(^\+?84\d{9,10}$)`)
	emailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	cmtPattern   = regexp.MustCompile(`(^\d{6}$)|(^\d{8,9}$)|(^(\w|\d){12}$)`)
	ccPattern    = regexp.MustCompile(`(^\d{12}$)`)
	hcPattern    = regexp.MustCompile(`(^\w\d+$)`)
)

func PhoneNumberValidator(phone *string) bool {
	if strings.HasPrefix(*phone, "+840") {
		*phone = strings.Replace(*phone, "+840", "0", 1)
	} else if strings.HasPrefix(*phone, "840") {
		*phone = strings.Replace(*phone, "840", "0", 1)
	} else if strings.HasPrefix(*phone, "+84") {
		*phone = strings.Replace(*phone, "+84", "0", 1)
	} else if strings.HasPrefix(*phone, "84") {
		*phone = strings.Replace(*phone, "84", "0", 1)
	}
	return phonePattern.MatchString(*phone)
}
func EmailValidate(email *string) bool {
	return emailPattern.MatchString(*email)
}
func IsCMT(identity string) bool {
	return cmtPattern.MatchString(identity) && !ccPattern.MatchString(identity)
}
func IsCC(identity string) bool {
	return ccPattern.MatchString(identity)
}
func IsHC(identity string) bool {
	return hcPattern.MatchString(identity) && !cmtPattern.MatchString(identity)
}

