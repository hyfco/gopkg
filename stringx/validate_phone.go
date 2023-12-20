package stringx

import (
	"errors"
	"regexp"
)

var (
	CHN              = `^1\d{10}$`
	TWN              = `^09\d{8}$`
	HKG              = `^[5689]\d{7}$`
	MAC              = `^6\d{7}$`
	JPN              = `^0\d{9,10}$`
	USA              = `^0\d{9,10}$`
	kOR1             = `^01[016789]\d{7}$`
	kOR2             = `^01[016789]\d{8}$`
	CAN              = `^[3456789]\d{9}$`
	regexPhoneChina  *regexp.Regexp
	regexPhoneTaiwan *regexp.Regexp
	regexPhoneHK     *regexp.Regexp
	regexPhoneMacao  *regexp.Regexp
	regexPhoneJPN    *regexp.Regexp
	regexPhoneUSA    *regexp.Regexp
	regexPhoneKOR1   *regexp.Regexp
	regexPhoneKOR2   *regexp.Regexp
)

func init() {
	regexPhoneChina, _ = regexp.Compile(CHN)
	regexPhoneTaiwan, _ = regexp.Compile(TWN)
	regexPhoneHK, _ = regexp.Compile(HKG)
	regexPhoneMacao, _ = regexp.Compile(MAC)
	regexPhoneJPN, _ = regexp.Compile(JPN)
	regexPhoneUSA, _ = regexp.Compile(USA)
	regexPhoneKOR1, _ = regexp.Compile(kOR1)
	regexPhoneKOR2, _ = regexp.Compile(kOR2)

}

func ValidateChinaPhoneNumber(phone string) error {
	if regexPhoneChina.Match([]byte(phone)) {
		return nil
	}
	return errors.New("invalid phone number")
}
