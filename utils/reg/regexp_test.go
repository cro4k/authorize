package reg

import "testing"

func TestRegexp(t *testing.T) {
	t.Log(Username.MatchString("123"))
	t.Log(Username.MatchString("123#"))
}
