package reg

import "regexp"

var (
	Username, _  = regexp.Compile("^[a-zA-Z\\d]{3}\\w{0,27}$")
	Cellphone, _ = regexp.Compile("^\\d{11}$")
	Email, _     = regexp.Compile("^[\\w.-]+@[\\w.-]+$")

	Number, _    = regexp.Compile("^\\d+$")
	Capital, _   = regexp.Compile("^[A-Z]+$")
	Minuscule, _ = regexp.Compile("^[a-z]+$")
	Symbol, _    = regexp.Compile("^[{}:\"|<>?,./;'\\\\[\\]~@#$%^&*()_+=-]+$")
	Letter, _    = regexp.Compile("^[A-Za-z]+$")

	MD5, _ = regexp.Compile("^[a-fA-F\\d]{32}$")
)

func SafePassword(pwd string) bool {
	var safe int
	if Number.FindString(pwd) != "" {
		safe++
	}
	if Capital.FindString(pwd) != "" {
		safe++
	}
	if Minuscule.FindString(pwd) != "" {
		safe++
	}
	if Symbol.FindString(pwd) != "" {
		safe++
	}
	return len(pwd) > 7 && safe > 2
}
