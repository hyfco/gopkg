package stringx

import (
	"errors"
	"regexp"
)

// 包含字母数字下划线
var regexNameStr = `^\w+$`

var regexName *regexp.Regexp

func init() {
	regexName = regexp.MustCompile(regexNameStr)
}

var (
	ErrNameInvalidate = errors.New("name invalidate")
)

func ValidateName(src string) error {
	if regexName.MatchString(src) {
		return nil
	}

	return ErrNameInvalidate
}
