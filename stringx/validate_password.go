package stringx

import (
	"errors"
	"regexp"
)

// 字母开头，只能包含字母数字下划线
var regexPasswordStr = `^[a-zA-Z]\w{5,17}$`

var regexPassword *regexp.Regexp

func init() {
	regexPassword = regexp.MustCompilePOSIX(regexEmailStr)
}

var (
	ErrPasswordInvalidate = errors.New("password  invalidate failed.")
)

func ValidatePassword(src string) error {

	if regexPassword.MatchString(src) {
		return nil
	}

	return ErrPasswordInvalidate
}
