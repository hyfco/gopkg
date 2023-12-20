package stringx

import "testing"

func TestValidateEmail(t *testing.T) {
	if err := ValidateEmail("hello@qq.com"); err != nil {
		t.Error(err)
	}
}
