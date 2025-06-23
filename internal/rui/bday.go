package rui

import "time"

// IsBirthday returns true if it's Rui's birthday (June 24rd) today.
func IsBirthday() bool {
	t := time.Now()
	return t.Day() == 24 && t.Month() == time.June // June 24rd = Rui's birthday
}
