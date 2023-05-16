package etc

import (
	"unicode"
)

func IsValidPhoneNubmer(phone string) bool {
	if string(phone[:3]) != "998" || len(phone) != 12 {
		return false
	}
	for i := 3; i < len(phone); i++ {
		if !unicode.IsDigit(rune(phone[i])) {
			return false
		}
	}
	return true
}
