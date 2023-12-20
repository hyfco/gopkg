package stringx

import (
	"errors"
	"regexp"
	"strings"
)

var regexEmailStr = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
var regexEmail *regexp.Regexp

func init() {
	regexEmail = regexp.MustCompilePOSIX(regexEmailStr)
}

var (
	ErrEmailInvalidate = errors.New("email address invalidate")
)

func ValidateEmail(email string) error {
	tmp := strings.ToLower(email)
	if regexEmail.MatchString(tmp) {
		return nil
	}

	return ErrEmailInvalidate
}
